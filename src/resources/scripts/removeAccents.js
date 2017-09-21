const removeAccents = str => [
  { chr: 'a', regex: /[ÀÁÂÃÄȂàáâãäª]/ },
  { chr: 'e', regex: /[ÉÊËéêë]/ },
  { chr: 'i', regex: /[ÍÎÏíîï]/ },
  { chr: 'o', regex: /[ÓÔÕÖóôõöº]/ },
  { chr: 'u', regex: /[ÚÛÜúûü]/ },
  { chr: 'c', regex: /[ÇḈç]/ },
  { chr: 'n', regex: /[Ññ]/ }
].reduce((acum, accent) =>
  acum.replace(accent.regex, accent.chr),
  str
)
