#!/usr/bin/mawk -f
#/root/mawk/bin/mawk -f

# $ cd /mnt/c/www/node/pussybnb/sh
# $ $HOME/mawk/bin/mawk -f parser.awk old.csv new.csv
# $ ./parser.awk src/pornhub.com-db.csv > src/out.csv
# $ head -n 10 src/out.csv
# $ zcat src/pornhub.com-db-1478439787.zip | ./parser.awk > src/out.csv
# $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | ./parser.awk > src/out.csv
# $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | grep -oPa '(?<=pornhub\.com\/embed\/)([^\s]+)(?=\")'
# $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | grep -oEa '(?:pornhub\.com\/embed\/)([^\s]+)(?=\")'
# $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | ./parser.awk
# WORKS: $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | mawk '{gsub(/.*pornhub.com\/embed\/|".*/,"")}1'
# WORKS: $ zcat src/pornhub.com-db-1478439787.zip | head -n 3 | mawk '{gsub(/.*pornhub.com\/embed\/|".*/,"",$1);print $1}'
# Reading between lines:
# $ time sed -n 1100000,1100010p src/out.csv

# pm55367bdf99a19 => Premium videos

{
  Embed = $1;
  Thumbnail = $2;
  Flipbook = $3;
  Title = $4;
  Tags = $5;
  Categories = $6;
  Pornstars = $7;
  Duration = $8;
  Rating = $9;
  Upvote = $10;
  Downvote = $11;
}

BEGIN {
  FS="|";
}

# POSIX ERE (Extended Regular Expressions) (egrep)
# http://www.regular-expressions.info/posix.html
{
  # http://stackoverflow.com/a/40664660/1442219
  gsub(/.*pornhub\.com\/embed\/|".*/, "", Embed);
}

#match(Embed, /pornhub\.com\/embed\/[^"]*/) {
#  Id = substr(Embed, RSTART, RLENGTH);
#}

{
  #print Id"\t"Duration;
  print Embed"\t"Thumbnail"\t"Flipbook"\t"Title"\t"Tags"\t"Categories"\t"Pornstars"\t"Duration"\t"Rating"\t"Upvote"\t"Downvote;
}

# match csv by field
# http://stackoverflow.com/questions/28318219/matching-third-field-in-a-csv-with-pattern-file-in-gnu-linux-awk-sed-grep

# Difference:
# $ awk -F'|' 'NR == FNR {old[$2]; next} !($2 in old)' old.csv new.csv > out.csv
