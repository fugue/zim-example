const faker = require('faker')

exports.handler = async function(event) {
  return faker.helpers.createCard();
}
