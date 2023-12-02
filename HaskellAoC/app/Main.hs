module Main where
import DayOne
import DayTwo (dayTwoP1)
-------------------------------------------
main :: IO ()
main = do
    s <- readFile "./app/input.txt"
    print . dayTwoP1 $ lines s