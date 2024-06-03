grep $(sort unixbox/poem.txt | uniq -c | sort | head -1 | tail -c 5) unixbox/poem.txt
