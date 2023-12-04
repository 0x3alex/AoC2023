module Main where
import DayOne
import DayTwo (dayTwoP1, dayTwoP2)
import DayFour (dayFourP1)
-------------------------------------------
main :: IO ()
main = do
    s <- readFile "./app/input.txt"
    print $ dayFourP1 $ lines s