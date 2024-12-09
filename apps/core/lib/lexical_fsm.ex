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

  def from(self, from_state)
  when is_atom(from_state) do
    {self, {from_state}}
  end

  def any({self, {from_state}})
  when is_atom(from_state)
  do
    {self, {from_state, :any}}
  end

  def with({self, {from_state}}, characters)
  when is_atom(from_state)
  and is_binary(characters) do
    edge = {:with, String.to_charlist(characters)}

    {self, {from_state, edge}}
  end

  def with_range({self, {from_state}}, start_character, end_character)
  when is_atom(from_state)
  and is_integer(start_character)
  and is_integer(end_character) do
    edge = {:range, start_character, end_character}

    {self, {from_state, edge}}
  end

  def without({self, {from_state}}, characters)
  when is_atom(from_state)
  and is_binary(characters) do
    edge = {:without, String.to_charlist(characters)}

    {self, {from_state, edge}}
  end

  def without_range({self, {from_state}}, start_character, end_character)
  when is_atom(from_state)
  and is_integer(start_character)
  and is_integer(end_character) do
    edge = {:without_range, start_character, end_character}

    {self, {from_state, edge}}
  end

  def to({self = %__MODULE__{}, {from_state, edge}}, to_state)
  when is_atom(from_state)
  and (is_tuple(edge) or is_atom(edge))
  and is_atom(to_state)
  and not is_nil(to_state) do
    new_transitions = {from_state, edge, to_state}

    %__MODULE__{self | transitions: [new_transitions | self.transitions]}
  end

  def build(self = %__MODULE__{}) do
    reversed_transitions = Enum.reverse(self.transitions)

    %__MODULE__{self | built: true, transitions: reversed_transitions}
  end

  def parse(self = %__MODULE__{}, stream) do
    if self.built == false do
      raise "The lexical FSM needs to be built before being used."
    end

    parse(self, stream, [nil])
  end

  defp parse(self = %__MODULE__{}, (<<character::utf8>> <> rest_stream) = stream, [current_state | _] = states) do
    case accept(self.transitions, character, current_state) do
      nil -> {stream, states}
      new_state -> parse(self, rest_stream, [new_state | states])
    end
  end

  defp accept(transitions, character, current_state) do
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
