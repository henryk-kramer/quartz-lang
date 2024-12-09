defmodule QuartzLang.Core.String do
  alias QuartzLang.Core.LexicalFsm, as: Fsm

  def parse(stream) do
    fsm =
      Fsm.init()
      |> Fsm.finals([:closed])
      |> Fsm.nomatches([])

      |> Fsm.from(nil)      |> Fsm.with("\"")   |> Fsm.to(:opened)

      |> Fsm.from(:opened)  |> Fsm.with("\\")   |> Fsm.to(:escaped)
      |> Fsm.from(:opened)  |> Fsm.with("\"")   |> Fsm.to(:closed)
      |> Fsm.from(:opened)  |> Fsm.any()        |> Fsm.to(:content)

      |> Fsm.from(:content) |> Fsm.with("\\")   |> Fsm.to(:escaped)
      |> Fsm.from(:content) |> Fsm.with("\"")   |> Fsm.to(:closed)
      |> Fsm.from(:content) |> Fsm.any()        |> Fsm.to(:content)

      |> Fsm.from(:escaped) |> Fsm.any()        |> Fsm.to(:content)

      |> Fsm.build()

    # (state_history, value, position)
    # (state_history, value, literal, position)
    Fsm.parse(fsm, stream)
  end
end
