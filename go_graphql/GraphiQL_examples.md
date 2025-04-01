Запросы и мутации


```
mutation CreatePost {
  CreatePost(
    input: {
      Title: "How to create new GraphQL app"
      Content: "Мы создаем новый вид API с помощью GraphQL"
      Author: "User"
      Hero: "User picture link"
    }
  ) {
    id
    Title
    Author
  }
}

query GetOnePost {
  GetOnePost(id: "67ebe81b20c1580ed44c0031") {
    id
    Title
    Content
    Author
    Hero
    Published_At
    Updated_At
  }
}
```