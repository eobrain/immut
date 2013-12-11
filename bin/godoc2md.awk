{
  if( /^[A-Z][A-Z][A-Z]*$/ ) {
    print $0 "\n-----";
  } else {
    print;
  }
}
