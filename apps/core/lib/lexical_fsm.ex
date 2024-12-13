defmodule QuartzLang.Core.LexicalFsm do
  defstruct [
    final_states: [],
    nomatch_states: [],
    transitions: [],
    built: false
  ]

  def init(), do: %__MODULE__{}

  def finals(self, final_states) do
    %__MODULE__{self | final_states: final_states}
  end

  def nomatches(self, nomatch_states) do
    %__MODULE__{self | nomatch_states: nomatch_states}
  end

  def any(self, from_state, to_state) do
    transition = {:any, from_state, to_state}

    %__MODULE__{self | transitions: [transition | self.transitions]}
  end

  def with(self, from_state, characters, to_state) do
    transition = {:with, from_state, to_state, String.to_charlist(characters)}

    %__MODULE__{self | transitions: [transition | self.transitions]}
  end

  def with_range(self, from_state, {start_character, end_character}, to_state) do
    transition = {:with_range, from_state, to_state, start_character, end_character}

    %__MODULE__{self | transitions: [transition | self.transitions]}
  end

  def without(self, from_state, characters, to_state) do
    transition = {:without, from_state, to_state, String.to_charlist(characters)}

    %__MODULE__{self | transitions: [transition | self.transitions]}
  end

  def without_range(self, from_state, {start_character, end_character}, to_state) do
    transition = {:without_range, from_state, to_state, start_character, end_character}

    %__MODULE__{self | transitions: [transition | self.transitions]}
  end

  def build(self) do
    reversed_transitions = Enum.reverse(self.transitions)

    %__MODULE__{self | built: true, transitions: reversed_transitions}
  end

  def parse(self, stream) do
    if self.built == false do
      raise "The lexical FSM needs to be built before being used."
    end

    Enum.reduce_while(String.to_charlist(stream), nil, fn char, current_state ->
      next_state = accept(self.transitions, char, current_state)

      if is_nil(next_state) do
        {:halt, current_state}
      else
        {:cont, next_state}
      end
    end)
  end

  defp accept(transitions, character, current_state) do
    # Need to adapt pattern matching to {:any, from, to} etc.

    Enum.find_value(transitions, fn
      {^current_state, :any, to_state} ->
        to_state

      {^current_state, {:with, characters}, to_state} ->
        if Enum.member?(characters, character), do: to_state

      {^current_state, {:with_range, start_character, end_character}, to_state}
      when start_character <= character and character <= end_character  ->
        to_state

      {^current_state, {:with, characters}, to_state} ->
        if Enum.member?(characters, character), do: to_state

      {^current_state, {:with_range, start_character, end_character}, to_state}
      when not (start_character <= character and character <= end_character)  ->
        to_state

      _ -> nil
    end)
  end
end
