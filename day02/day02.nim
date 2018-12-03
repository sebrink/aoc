import strutils, sequtils, sugar, os, sets, tables

proc partOne(file: seq[string]): int =
    var restab = initTable[int, int]()
    var checksum = 0
    for line in file:
        var iTable = initTable[char,int]()
        for i in line:
            if iTable.hasKey(i):
                iTable[i] += 1
            else:
                iTable[i] = 1
        var interSet = initSet[int]()
        for x, y in iTable:
            interSet.incl y
        for y in interset:
            if y != 1:
                if restab.hasKey(y):
                    restab[y] += 1
                else:
                    restab[y] = 1
    result = 1
    for x, y in restab:
        result *= y

proc partTwo(file: seq[string]): string =
    #[let llen = file[0].len
    for xl in file:
        for yl in file:
            var diff = ""
            var wrong = 0
            for i in 0 ..< llen:
                if xl[i] == yl[i]:
                    diff.add xl[i]
                else:
                    inc wrong
                    if wrong > 1:
                        break
            if wrong == 1:
                return diff]#
  
    let lstLen = file[0].len
    for i in file:
        for j in file:
            var dif = ""
            var wrong = 0
            for k in 0 ..< lstLen:
                if i[k] == j[k]:
                    dif.add i[k]
                else:
                    inc wrong
                    if wrong > 1:
                       break
            if wrong == 1:
                return dif


let file = readFile("input.txt").splitLines.filterIt(it.len > 0)
echo file.partOne
echo file.partTwo