defmodule StringTest do
  use ExUnit.Case

  test "normal string" do
    QuartzLang.Core.String.parse("normal String")
  end
end
