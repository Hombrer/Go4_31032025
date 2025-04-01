# Manual(Создание приложения)

### Step 1
```
user@usercomp:~/Projects/Go4/go_graphql$ go run github.com/99designs/gqlgen init
Creating gqlgen.yml
Creating graph/schema.graphqls
Creating server.go
Generating...

Exec "go run ./server.go" to start GraphQL server
```

Files:  
 - *gqlgen.yml* - contains the gqlgen configurations.  
 - *server.go* - the application entry point that serves your GraphQL endpoint.
 - *graph/resolver.go* - contains the type for your resolvers.
 - *graph/schema.graphqls* - a file to write down your API schemas.
 - *graph/schema.resolvers.go* - contains the generated resolvers methods that you use to implement your API Mutation and Query types.
 - *graph/model/models_gen.go* - contains structs generated from the schema file.
 - *graph/generated/generated.go* - contains the generated runtime for GraphQL.

### Step 2
Setting up GraphQL Schema for Go
```
scalar Time

type Post {
  _id: String!
  Title: String!
  Content: String!
  Author: String!
  Hero: String!
  Published_At: Time!
  Updated_At: Time!
}
 
type DeletePostResponse {
  deletePostId: String!
}

type Query {
  GetAllPosts: [Post!]!
  GetOnePost(id: String!): Post!
}
 
input NewPost {
  Title: String!
  Content: String!
  Author: String
  Hero: String
  Published_At: Time
  Updated_At: Time
}
 
type Mutation {
  CreatePost(input: NewPost!): Post!
  UpdatePost(id: String!, input: NewPost): Post!
  DeletePost(id: String!): DeletePostResponse!
}

```

### Step 3 generate
```
go run github.com/99designs/gqlgen generate
```

### Step4 Setting Up the Database Dependencies

```
go get go.mongodb.org/mongo-driver/mongo
```

Добавить код для соединения с БД

### Step5 reload (для правильной работы нужно удалить tools.go)
```
go install github.com/air-verse/air@latest
air init
air
```

### Step 6 Add resolvers 
Добавить код в `graph/schema.resolvers.go`

### Step 7 Run server 
Запустить сервер: `go run ./server.go`   
Запустить браузер по адресу: http://localhost:8080/


