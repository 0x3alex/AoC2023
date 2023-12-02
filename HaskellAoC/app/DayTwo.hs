module DayTwo where
import Data.List.Split (splitOn, split)
import Data.Char (isDigit)
import Data.List (find, isInfixOf, elemIndex)
import Data.Maybe (fromMaybe)

dayTwoP1_4 :: [String] -> Bool
dayTwoP1_4 x = g "green" 13 && g "blue" 14 && g "red" 12
        where
                g j k= let idx = f j in ((idx == (-1)) || ((read (x !! (idx -1)) :: Int) <= k))
                f k = fromMaybe (-1) $ k `elemIndex` x

dayTwoP1_3 :: [String] -> Bool
dayTwoP1_3 = foldl (\acc x -> acc && dayTwoP1_4 (splitOn " " x)) True

dayTwoP1_2:: [String] -> Bool
dayTwoP1_2 = foldl (\acc x -> acc && dayTwoP1_3 (splitOn "," x)) True

dayTwoP1_1 :: String -> String -> Int
dayTwoP1_1 xs g = if dayTwoP1_2 (splitOn ";" xs) then read (filter isDigit g):: Int else 0

dayTwoP1:: [String] -> Int
dayTwoP1 = foldl (\acc x -> acc + dayTwoP1_1 (last $ g x) (head $ g x)) 0
        where g = splitOn ":"