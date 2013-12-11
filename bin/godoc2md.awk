BEGIN {
 code=0
}
{
  if( /^[A-Z][A-Z][A-Z]*$/ ) {
    if (code) {
      print "````\n";
      code = 0;
    }
    print
  } else {
    if (!code) {
      print "````go\n";
      code = 1;
    }
    print
  }
}
END {
    if (code) {
      print "````\n";
    }
}
