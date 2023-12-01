module Main where
import DayOne
-------------------------------------------
main :: IO ()
main = do
    s <- readFile "./app/input.txt"
    print . dayOneP2 . lines $ s