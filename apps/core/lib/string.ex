defmodule QuartzLang.Core.String do
  alias QuartzLang.Core.LexicalFsm, as: Fsm

  def parse(stream) do
    fsm =
      Fsm.init()
      |> Fsm.finals([:closed])
      |> Fsm.nomatches([])

      |> Fsm.with(nil,        "\"",   :opened)
      |> Fsm.with(:opened,    "\\",   :escaped)
      |> Fsm.with(:opened,    "\"",   :closed)
      |> Fsm.any( :opened,            :content)
      |> Fsm.with(:content,   "\\",   :escaped)
      |> Fsm.with(:content,   "\"",   :closed)
      |> Fsm.any( :content,           :content)
      |> Fsm.any( :escaped,           :content)

      |> Fsm.build()

    # (state_history, value, position)
    # (state_history, value, literal, position)
    Fsm.parse(fsm, stream)
  end
end
