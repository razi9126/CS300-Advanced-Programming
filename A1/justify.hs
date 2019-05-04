-- ========================================================================================================================== --
-- MOHAMMAD RAZI UL HAQ 2020-10-0015

-- README
-- I really don't think there are any bugs here.
-- I tried to account for most of the cases but still
-- if you do find some erroneous cases please do inform
-- me as well so that I can look into them and avoid 
-- making those in the future. 
-- The code however, doesn't give correct output for the
-- last line of the test paragraph ("the past."). It inserts
-- spaces in this line too. By now am too tired to figure out why it
-- is doing this. Happy checking :)

--
--                                                          ASSIGNMENT 1
--
--      A common type of text alignment in print media is "justification", where the spaces between words, are stretched or
--      compressed to align both the left and right ends of each line of text. In this problem we'll be implementing a text
--      justification function for a monospaced terminal output (i.e. fixed width font where every letter has the same width).
--
--      Alignment is achieved by inserting blanks and hyphenating the words. For example, given a text:
--
--              "He who controls the past controls the future. He who controls the present controls the past."
--
--      we want to be able to align it like this (to a width of say, 15 columns):
--
--              He who controls
--              the  past cont-
--              rols  the futu-
--              re. He  who co-
--              ntrols the pre-
--              sent   controls
--              the past.
--


-- ========================================================================================================================== --


import Data.List
import Data.Char

text1 = "He who controls the past controls the future. He who controls the present controls the past."
text2 = "A creative man is motivated by the desire to achieve, not by the desire to beat others."


-- ========================================================================================================================== --



-- MOHAMMAD RAZI UL HAQ 2020-10-0015



-- ========================================================= PART 1 ========================================================= --


--
-- Define a function that splits a list of words into two lists, such that the first list does not exceed a given line width.
-- The function should take an integer and a list of words as input, and return a pair of lists.
-- Make sure that spaces between words are counted in the line width.
--
-- Example:
--    splitLine ["A", "creative", "man"] 12   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 11   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 10   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 9    ==>   (["A"], ["creative", "man"])
--


splitLine :: [String] -> Int -> ([String], [String])
-- Function definition here
wordLengths []     = 0
wordLengths (x:xs) = length x + 1 + wordLengths xs

-- Takes a empty list tuple initially and adds values to it until it reaches the length n
sLine [] _ (a,b) = (a,b)
sLine (x:xs) n (a,b)
	| length a == 0, length x > n = (a,b++[x]++xs)
	| wordLengths a + length x > n = (a,b++[x]++xs)
	| otherwise = sLine (xs) n (a++[x],b)



splitLine [] _ = ([],[])
splitLine (x:xs) n = sLine (x:xs) n ([],[])


-- ========================================================= PART 2 ========================================================= --


--
-- To be able to align the lines nicely. we have to be able to hyphenate long words. Although there are rules for hyphenation
-- for each language, we will take a simpler approach here and assume that there is a list of words and their proper hyphenation.
-- For example:

enHyp = [("creative", ["cr","ea","ti","ve"]), ("controls", ["co","nt","ro","ls"]), ("achieve", ["ach","ie","ve"]), ("future", ["fu","tu","re"]), ("present", ["pre","se","nt"]), ("motivated", ["mot","iv","at","ed"]), ("desire", ["de","si","re"]), ("others", ["ot","he","rs"])]


--
-- Define a function that splits a list of words into two lists in different ways. The first list should not exceed a given
-- line width, and may include a hyphenated part of a word at the end. You can use the splitLine function and then attempt
-- to breakup the next word with a given list of hyphenation rules. Include a breakup option in the output only if the line
-- width constraint is satisfied.
-- The function should take a hyphenation map, an integer line width and a list of words as input. Return pairs of lists as
-- in part 1.
--
-- Example:
--    lineBreaks enHyp 12 ["He", "who", "controls."]   ==>   [(["He","who"], ["controls."]), (["He","who","co-"], ["ntrols."]), (["He","who","cont-"], ["rols."])]
--
-- Make sure that words from the list are hyphenated even when they have a trailing punctuation (e.g. "controls.")
--
-- You might find 'map', 'find', 'isAlpha' and 'filter' useful.
--

-- HELPER FUNCTIONS
-- joins the text into the tuple
joinNew (a,b) text = dropSecond ( (a ++ [text ++ "-"],b)) text

-- alters b element of the tuple so that the head of b contains only the remaining letters of the hyphenated word
dropSecond (a,[]) _ = (a, [])
dropSecond (a,b) text =(a,(head b \\ text) :tail b)

-- returns each possible hyphenation
sum' (x1:x2:xs)= [x1] ++ sum'((x1++x2):xs)
sum' (x1:xs)= [x1]
sum' []=[]

-- Finds the second part from the tuple list
search' "" _ = []
search' _ [] = []
search' x ((a,b):xs) = if x == a then b else search' x xs

-- Returns the head of list b
f [] _ = []
f (x:xs) n 
	| snd(splitLine (x:xs) n) == [] = ""
	| otherwise = head(snd(splitLine (x:xs) n))

-- Deletes the empty element tuples
del deleted xs = [ x | x <- xs, x /= deleted ]

otherValues (a,b) n m z
	| wordLengths (fst (joinNew (a,b) z)) - 1 > n = ([],[])
	| otherwise = joinNew (a,b) z


lineBreaks :: [(String, [String])] -> Int -> [String] -> [([String], [String])]
-- Function definition here
lB m n (x:xs) = [(splitLine (x:xs) n)] ++ map (otherValues (splitLine (x:xs) n) n m) val where 
	val = sum' (search' (filter (isAlpha) (f (x:xs) n)) m)
lineBreaks m n (x:xs) = del (([],[])) (lB m n (x:xs))

-- ========================================================= PART 3 ========================================================= --



-- Define a function that inserts a given number of blanks (spaces) into a list of strings and outputs a list of all possible
-- insertions. Only insert blanks between strings and not at the beginning or end of the list (if there are less than two
-- strings in the list then return nothing). Remove duplicate lists from the output.
-- The function should take the number of blanks and the the list of strings as input and return a lists of strings.

-- Example:
--    blankInsertions 2 ["A", "creative", "man"]   ==>   [["A", " ", " ", "creative", "man"], ["A", " ", "creative", " ", "man"], ["A", "creative", " ", " ", "man"]]

-- Use let/in/where to make the code readable

-- Helper functions
 -- Inserts 1 space in the given list of strings
insSp :: [String] -> [[String]]
insSp [y,x] = [[y]++[" "]++  [x]]
insSp (x:xs) = [x:" ": xs] ++ (map (\y -> x:y) (insSp xs))

 -- Inserts 1 space in the given list of LIST of strings
listSp [x] = insSp x
listSp (x:xs) = (insSp x)  ++ (listSp xs)

-- Used these 2 from internet. Names are self explanatory
applyNTimes :: Int -> (a -> a) -> a -> a
applyNTimes n f x 
    | n == 0        = x
    | otherwise     = f (applyNTimes (n-1) f x)

rmdups :: Eq a => [a] -> [a]
rmdups [] = []
rmdups (x:xs)   | x `elem` xs   = rmdups xs
                | otherwise     = x : rmdups xs



blankInsertions :: Int -> [String] -> [[String]]
-- Function definition here
blankInsertions 0 list 	 = [list]
blankInsertions _ [x] 	 = []
blankInsertions 1 (list) = insSp list
blankInsertions n (list) = rmdups (applyNTimes (n-1) listSp (insSp list))





-- ========================================================= PART 4 ========================================================= --


--
-- Define a function to score a list of strings based on four factors:
--
--    blankCost: The cost of introducing each blank in the list
--    blankProxCost: The cost of having blanks close to each other
--    blankUnevenCost: The cost of having blanks spread unevenly
--    hypCost: The cost of hyphenating the last word in the list
--
-- The cost of a list of strings is computed simply as the weighted sum of the individual costs. The blankProxCost weight equals
-- the length of the list minus the average distance between blanks (0 if there are no blanks). The blankUnevenCost weight is
-- the variance of the distances between blanks.
--
-- The function should take a list of strings and return the line cost as a double
--
-- Example:
--    lineCost ["He", " ", " ", "who", "controls"]
--        ==>   blankCost * 2.0 + blankProxCost * (5 - average(1, 0, 2)) + blankUnevenCost * variance(1, 0, 2) + hypCost * 0.0
--        ==>   blankCost * 2.0 + blankProxCost * 4.0 + blankUnevenCost * 0.666...
--
-- Use let/in/where to make the code readable
--
-- Helper Functions
-- Calculates the list of values used for average and variance
numbSp [] list count = list ++ [count]
numbSp (x:xs) list count
	| x==" " = numbSp xs (list ++ [count]) 0
	| otherwise = numbSp xs list (count+1)


-- Tells whether there is a hyphenated word at the end of the list
isHyphen [] = 0.0
isHyphen list
	| last (last list) == '-' = 1.0
	| otherwise = 0.0

-- Calculates the number of spaces in a list
numberOfSpaces [] = 0
numberOfSpaces (x:xs) 
	| x==" " = 1 + (numberOfSpaces xs)
	| otherwise = 0 + (numberOfSpaces xs)

sumOfVal [] = 0
sumOfVal (x:xs) = x + sumOfVal xs

-- Calculates average
average [] = 0
average (x:xs) = fromIntegral(sumOfVal (x:xs)) / fromIntegral(length (x:xs))

-- Used variance and stdev from internet
var xs = (stdev xs) * (stdev xs)
   
stdev :: [Float] -> Float
stdev xs = sqrt . average . map ((^2) . (-) axs) $ xs
           where average = (/) <$> sum <*> realToFrac . length
                 axs     = average xs


---- Do not modify these in the submission ----
blankCost = 1.0
blankProxCost = 1.0
blankUnevenCost = 1.0
hypCost = 1.0
-----------------------------------------------


lineCost :: [String] -> Double
lineCost (x:xs) = 
	let a = fromIntegral(length (x:xs))
	in (blankCost * fromIntegral (numberOfSpaces (x:xs))) + (blankProxCost * (a - average (numbSp (x:xs) [] 0))) + (hypCost * isHyphen (x:xs)) + (blankUnevenCost * realToFrac (var (numbSp (x:xs) [] 0))) 



-- ========================================================= PART 5 ========================================================= --

--
-- Define a function that returns the best line break in a list of words given a cost function, a hyphenation map and the maximum
-- line width (the best line break is the one that minimizes the line cost of the broken list).
-- The function should take a cost function, a hyphenation map, the maximum line width and the list of strings to split and return
-- a pair of lists of strings as in part 1.
--
-- Example:
--    bestLineBreak lineCost enHyp 12 ["He", "who", "controls"]   ==>   (["He", "who", "cont-"], ["rols"])
--
-- Use let/in/where to make the code readable
--

-- Helper Functions

-- Simply joins the values to return a tuple
joinTup x y = (y,x)

-- Takes the line width and tuple of values returned from Linebreaks to insert appropriate number of blanks
sendToBI _ [] = []
sendToBI n ((a,b):xs) = (map (joinTup b) (blankInsertions (n - (wordLengths a) +1) a)) ++ (sendToBI n xs)

-- Takes the list of tuples and the cost function and returns the tuple with the minimum cost
sendToLC [(a,b)] _ = (a,b)
sendToLC ((a,b):(c,d):xs) function
	| function a < function c = sendToLC ((a,b):xs) function
	| otherwise = sendToLC ((c,d):xs) function


bestLineBreak :: ([String] -> Double) -> [(String, [String])] -> Int -> [String] -> ([String], [String])
-- Function definition here
bestLineBreak fun wordmap n (x:xs) = sendToLC (sendToBI n values) fun
	where values = lineBreaks wordmap n (x:xs)

--
-- Finally define a function that justifies a given text into a list of lines satisfying a given width constraint.
-- The function should take a cost function, hyphenation map, maximum line width, and a text string as input and return a list of
-- strings.
--
-- 'justifyText lineCost enHyp 15 text1' should give you the example at the start of the assignment.
--
-- You might find the words and unwords functions useful.
--



justifyText :: ([String] -> Double) -> [(String, [String])] -> Int -> String -> [String]
-- Function definition here
justifyText func1 wordmap n [] = []
justifyText func1 wordmap n dat = [unwords firstList] ++  justifyText func1 wordmap n (unwords (snd (bestLineBreak func1 wordmap n (words dat))) )
	where firstList = (fst (bestLineBreak func1 wordmap n (words dat)))














