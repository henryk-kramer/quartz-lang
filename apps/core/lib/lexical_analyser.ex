defmodule QuartzLang.Core.LexicalAnalyser do

  def run(text) do
    parse(text, [])
  end

  def parse("", tokens), do: Enum.reverse(tokens)
  def parse(text, tokens) do

  end

end
