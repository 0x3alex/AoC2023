module DayOne (dayOneP1, dayOneP2) where

import Data.Char (isDigit, digitToInt, intToDigit)
import Data.Maybe (isNothing, fromJust, fromMaybe)
import Data.List (isInfixOf)


stringToInt :: String -> Int
stringToInt = read

numNames = ["one","two","three","four","five","six","seven","eight","nine"]

lookUpList = zip numNames ['1'..'9']

lookUp :: String -> Char
lookUp s = fromMaybe ' ' $ lookup s lookUpList

f :: String -> String -> String
f [] _ = []
f k@(x:xs) a
    | isDigit x = x : f xs ""
    | not . null $ mp = (lookUp . head $ mp) : f k ""
    | otherwise = f xs (a++[x])
    where   m ks = filter (`isInfixOf` ks) numNames
            mp = m (a++[x])

dayOneP2 :: [String] -> Int
dayOneP2 = dayOneP1 . map (`f` "") 

dayOneP1 :: [String] -> Int
dayOneP1 = foldr (\ x -> (+) (buildDigits . extractDigits $ x)) 0
    where
        extractDigits = filter isDigit
        buildDigits s
            | length s == 1 = stringToInt (s ++ s)
            | length s > 2 = stringToInt (head s : [last s])
            | otherwise = stringToInt s
