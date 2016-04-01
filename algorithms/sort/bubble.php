<?php
function bubbleSort(array &$s)
{
    $swapCounter = 0;
    $len = count($s)-1;
    $count = 0;
    do {
        $count++;
        $swapped = false;
        print_r($s);
        echo 'swapCounter is '.$swapCounter.PHP_EOL;
        for ($i=0; $i < $len; $i++) {
            if ($s[$i] > $s[$i+1]) {
                $tmp = $s[$i+1];
                $s[$i+1] = $s[$i];
                $s[$i] = $tmp;
                $swapped = true;
                $swapCounter++;
            }
        }
        $len--;
    } while ($swapped);
    echo 'Loop counter is '.$count.PHP_EOL;
}
$s = [3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48];
print_r($s);
bubbleSort($s);
print_r($s);
