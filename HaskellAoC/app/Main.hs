module Main where
import DayOne
import DayTwo (dayTwoP1, dayTwoP2)
-------------------------------------------
main :: IO ()
main = do
    s <- readFile "./app/input.txt"
    print . dayTwoP2 $ lines s