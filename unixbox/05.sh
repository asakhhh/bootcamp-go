$a = grep $(sort poem.txt | uniq -c | sort | head -1 | tail -c 5) poem.txt
echo -n $a