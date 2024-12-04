defmodule QuartzLang.Core.LexicalFsm do
  defstruct [
    final_states: [],
    nomatch_states: [],
    transitions: []
  ]

  def init(), do: %__MODULE__{}

  def finals(self, final_states) do
    %__MODULE__{self | final_states: final_states}
  end

  def nomatches(self, nomatch_states) do
    %__MODULE__{self | nomatch_states: nomatch_states}
  end

  def from(self, from_state) do
    {self, {from_state}}
  end

  def with({self, {from_state}}, characters) do
    {self, {from_state, characters}}
  end

  def to({self, {from_state, characters}}, to_state) do
    new_transitions = assemble_transition(from_state, characters, to_state)
    old_transitions = self.transitions

    # Need to look at transitions in reverse because of defined precedence
    %__MODULE__{self | transitions: [new_transitions | old_transitions]}
  end

  defp assemble_transition(from_states, characters, to_states) do

  end
end
