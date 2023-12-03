module Main where
import DayOne
import DayTwo (dayTwoP1, dayTwoP2)
import DayThree (dayThreeP1)
-------------------------------------------
main :: IO ()
main = do
    s <- readFile "./app/input.txt"
    print $ dayThreeP1 (lines s) 0 0