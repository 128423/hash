db.users.drop();
db.products.drop();

db.users.insertMany([
  {
    _id: 1,
    first_name: "joao",
    last_name: "da silva",
    date_of_birth: new Date("2016-05-18T16:00:00Z")
  },
  {
    _id: 2,
    first_name: "janaina",
    last_name: "da silva",
    date_of_birth: new Date()
  }
]);


db.products.insertMany([
  {
    _id: 1,
    price_in_cents: 100,
    title: "Carro",
    description: "Marea turbo, ta pegando fogo meu",
  },
  {
    _id: 2,
    price_in_cents: 1000,
    title: "Bomba",
    description: "Fiat 147, vai quebrar",
  }
]);
