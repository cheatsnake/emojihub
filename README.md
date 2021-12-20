
<a href="https://emojihub.herokuapp.com/"><img src="https://i.ibb.co/Cht49d8/Screenshot-15.jpg" alt="EmojiHub" border="0" style="{margin: 0 auto; width: 100%;}" /></a>

## :mag: Overview
EmojiHub provides an opportunity to get random emojis from already sorted categories and groups. You can also get a whole list of emojis by a certain category, group, or get the entire emoji database consisting of 1629 objects. All emoji data is stored in a simple JSON object from which you can get html codes to insert into your web applications.

## API Documentation
The service has two main endpoints:
- To get random emojis
```
https://emojihub.herokuapp.com/api/random
```
```js
{
  name: "hugging face",
  category: "smileys and people",
  group: "face positive",
  htmlCode: [
    "&#129303;"
  ],
  unicode: [
    "U+1F917"
  ]
}
```
- To get an array of all emojis
```
https://emojihub.herokuapp.com/api/all
```
Each endpoint can be supplemented by a path to a specific category or group
```
/category_{category_name}
```
```
/group_{group_name}
```

Table with all available emoji categories and groups
| Category           | Groups                                                                                                                                             |
| ------------------ |:-------------------------------------------------------------------------------------------------------------------------------------------------- |
| smileys_and_people | body<br>cat_face<br>clothing<br>creature_face<br>emotion<br>face_negative<br>face_neutral<br>face_positive<br>face_positive<br>face_role<br>face_sick<br>family<br>monkey_face<br>person<br>person_activity<br>person_gesture<br>person_role<br>skin_tone |
| animals_and_nature | animal_amphibian<br>animal_bird<br>animal_bug<br>animal_mammal<br>animal_marine<br>animal_reptile<br>plant_flower<br>plant_other                                                                                                                                                |
| food_and_drink     | dishware<br>drink<br>food_asian<br>food_fruit<br>food_prepared<br>food_sweat<br>food_vegetable                                                                                                                                                    |
| travel_and_places  | travel_and_places                                                                                                                                                   |
| activities         | activities                                                                                                                                                   |
| objects            | objects                                                                                                                                                   |
| symbols            | symbols                                                                                                                                                   |
| flags              | flags  
