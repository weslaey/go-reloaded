# go-reloaded

Petit projet en Go pour faire du traitement et de la modification de texte automatique via des balises. Le programme lit un fichier source (sample.txt), applique les règles demandées, et balance le résultat propre dans result.txt.

## Ce que le programme gère :
- La casse : (up), (low), (cap) pour modifier le mot juste avant. On peut aussi cibler plusieurs mots d'un coup avec un nombre, genre (up, 2).
- Les conversions : (hex) et (bin) pour transformer des nombres écrits en hexa ou binaire directement en décimal.
- Les articles : Il gère tout seul le passage de a à an quand le mot d'après commence par une voyelle ou un h.
- La ponctuation : Un petit coup de propre sur les espaces autour de la ponctuation (., ,, !, ?, etc.) pour que ce soit nickel.
