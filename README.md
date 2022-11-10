# EmojiHub

😎 EmojiHub provides an opportunity to get random emojis from already sorted categories and groups. You can also get a whole list of emojis by a certain category, group, or get the entire emoji database consisting of 1791 objects. All emoji data is stored in a simple JSON object from which you can get html codes to insert into your web applications.

## 📄 API documentation

-   Get random emoji

```rs
GET /api/random
```

-   Get an array of all emojis

```rs
GET /api/all
```

Each endpoint can be supplemented by a path to a specific category or group:

```
/category/{category-name}
```

```
/group/{group-name}
```

📚 Table with all available emoji categories and groups:

| Category           | Groups                                                                                                                                                                                                                  |
| ------------------ | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| smileys-and-people | body, cat-face, clothing, creature-face, emotion, face-negative, face-neutral, face-positive, face-positive, face-role, face-sick, family, monkey-face, person, person-activity, person-gesture, person-role, skin-tone |
| animals-and-nature | animal-amphibian, animal-bird, animal-bug, animal-mammal, animal-marine, animal-reptile, plant-flower, plant-other                                                                                                      |
| food-and-drink     | dishware, drink, food-asian, food-fruit, food-prepared, food-sweat, food-vegetable                                                                                                                                      |
| travel-and-places  | travel-and-places                                                                                                                                                                                                       |
| activities         | activities                                                                                                                                                                                                              |
| objects            | objects                                                                                                                                                                                                                 |
| symbols            | symbols                                                                                                                                                                                                                 |
| flags              | flags                                                                                                                                                                                                                   |

### 🎯 Examples

```
api/random/group/face-positive
```

```
api/random/category/food-and-drink
```

```
api/all/category/travel-and-places
```

```
api/all/group/animal-bird
```

## 🚀 Server startup

1. Clone this repository:

```sh
git clone -b go-version --single-branch https://github.com/cheatsnake/emojihub.git
```

2. Inside the project, run this command to install the necessary packages:

```sh
go mod download
```

> Make sure you have already [installed Go](https://go.dev) on your computer.

3. Start the server by running the last command:

```sh
go run cmd/main.go
```

> The server will start at the address: http://localhost:4000

## 🐳 Docker container startup

Run the following commands to create an image and start the container:

```sh
docker build -t emojihub . --target production
```

```sh
docker run -p 4000:4000 --name emojihub-server emojihub
```
