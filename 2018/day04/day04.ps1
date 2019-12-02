# Inspiried by a previous answer on reddit, not claiming as my own code
$data = Get-Content '.\input.txt' | Sort-Object

$guards = @{}
$curDate = ''
$curGuard = -1
[int]$sleepMinute = -1
foreach ($line in $data) {
    $tokens = $line.Replace('[', '').Replace(']', '').Replace(':', ' ').Replace('#', '').Split(' ')
    if ($tokens[0] -ne $curDate -or [int]$tokens[1] -eq 23) {
        $curDate = $tokens[0]
        if ($tokens[4] -ne 'asleep') {
            [int]$curGuard = $tokens[4]
        }
        if ($null -eq $guards[$curGuard]) {
            $guards[$curGuard] = New-Object int[] 60
        }
    }
    if ($tokens[3] -eq 'falls') {
        [int]$sleepMinute = $tokens[2]
    }
    elseif ($tokens[3] -eq 'wakes') {
        for ($i = $sleepMinute; $i -lt [int]$tokens[2]; $i++) {
            $guards[$curGuard][$i]++
        }
    }
}
$curBestGuard = -1
$curBestGuardTime = -1
$curBestMinute = -1
foreach ($g in $guards.Keys) {
    $sum = 0
    $guards[$g] | ForEach-Object {$sum += $_}
    if ($sum -gt $curBestGuardTime) {
        $curBestGuard = $g
        $curBestGuardTime = $sum
        $maximum = ($guards[$g] | Measure-Object -Max).Maximum
        [int]$curBestMinute = [Array]::IndexOf($guards[$g], [int]$maximum) 
        
    }
}

$check = [int]$curBestGuard * [int]$curBestMinute
Write-Host $check

$curBestGuard = -1
$curBestGuardTime = -1
$curBestMinute = -1

foreach ($g in $guards.Keys) {
    $maximum = ($guards[$g] | Measure-Object -Max).Maximum
    if ($maximum -gt $curBestMinute) {
        [int]$curBestGuardTime = [Array]::IndexOf($guards[$g], [int]$maximum)
        $curBestGuard = $g
        $curBestMinute = $maximum
    }
}
$check = [int]$curBestGuard * [int]$curBestGuardTime
Write-Host $check