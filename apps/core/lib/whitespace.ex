defmodule QuartzLang.Core.Whitespace do
  alias QuartzLang.Core.LexicalFsm, as: Fsm

  def parse(stream) do
    fsm =
      Fsm.init()
      |> Fsm.finals([:whitespace])
      |> Fsm.nomatches([:none])
      |> Fsm.from(nil)          |> Fsm.with("\t\s") |> Fsm.to(:whitespace)
      |> Fsm.from(:whitespace)  |> Fsm.with("\t\s") |> Fsm.to(:whitespace)
      |> Fsm.build()

    Fsm.parse(fsm, stream)
  end
end
