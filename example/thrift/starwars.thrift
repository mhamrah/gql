service Query {
    Character Hero (1: required Episode episode)
    Human Human(1: IdRequest request)
    Droid Droid(1: IdRequest request)
}

enum Episode {
    NEWHOPE,
    EMPIRE,
    JEDI
}

struct IdRequest {
    1: required string id;
}

struct Human {
    1: required string id;
    2: optional string name;
    3: optional list<Character> friends;
    4: optional list<Episode> appearsIn;
    5: optional string home_planet;
}

struct Character {
    1: required string id;
    2: optional string name;
    3: optional list<Character> friends;
    4: optional list<Episode> appearsIn;
}
struct Droid {
    1: required string id;
    2: optional string name;
    3: optional list<Character> friends;
    4: optional list<Episode> appearsIn;
    5: optional string primary_function;
}