defmodule QuartzLang.Core.Whitespace do
  alias QuartzLang.Core.LexicalFsm, as: Fsm

  def parse(stream) do
    fsm =


    fsm =
      Fsm.finals([:whitespace]).nomatches([:none])
      |> Fsm.from(:none, :whitespace).with("\t\s").to([:whitespace])

    Fsm.parse(stream)
  end
end
