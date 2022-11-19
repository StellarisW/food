/*
 Navicat Premium Data Transfer

 Source Server Type    : MongoDB
 Source Server Version : 40223
 Source Schema         : food

 Target Server Type    : MongoDB
 Target Server Version : 40223
 File Encoding         : 65001

 Date: 03/10/2022 22:12:16
*/


// ----------------------------
// Collection structure for recipe
// ----------------------------
db.getCollection("recipe").drop();
db.createCollection("recipe");
db.getCollection("recipe").createIndex({
    "recipe_id": NumberInt("-1")
}, {
    name: "recipe_id_NumberInt(\"-1\")",
    unique: true
});

// ----------------------------
// Documents of recipe
// ----------------------------
db.getCollection("recipe").insert([ {
    _id: ObjectId("6336eb430d300000ef083348"),
    "recipe_id": NumberLong("38"),
    name: "Low-Fat Berry Blue Frozen Dessert",
    "cook_time": NumberLong("86400"),
    "perp_time": NumberLong("2700"),
    "total_time": NumberLong("89100"),
    description: "Make and share this Low-Fat Berry Blue Frozen Dessert recipe from Food.com.",
    images: [
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/YUeirxMLQaeE1h3v3qnM_229%20berry%20blue%20frzn%20dess.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/AFPDDHATWzQ0b1CDpDAT_255%20berry%20blue%20frzn%20dess.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/UYgf9nwMT2SGGJCuzILO_228%20berry%20blue%20frzn%20dess.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/PeBMJN2TGSaYks2759BA_20140722_202142.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/picuaETeN.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/38/pictzvxW5.jpg"
    ],
    category: "Frozen Desserts",
    keywords: [
        "Dessert",
        "Low Protein",
        "Low Cholesterol",
        "Healthy",
        "Free Of...",
        "Summer",
        "Weeknignt",
        "Freezer",
        "Easy"
    ],
    ingredients: [
        "blueberries",
        "granulated sugar",
        "vanilla yogurt",
        "lemon juice"
    ],
    calories: 170.9,
    fat: 2.5,
    "saturated_fat": 1.3,
    cholesterol: 8,
    sodium: 29.8,
    carbonhydrate: 37.1,
    fiber: 3.6,
    sugar: 30.2,
    protein: 3.2,
    instruction: [
        "Toss 2 cups berries with sugar",
        "Let stand for 45 minutes, stirring occasionally.",
        "Transfer berry-sugar mixture to food processor.",
        "Add yogurt and process until smooth.",
        "Strain through fine sieve. Pour into baking pan (or transfer to ice cream maker and process according to manufacturers' directions). Freeze uncovered until edges are solid but centre is soft.  Transfer to processor and blend until smooth again.",
        "Return to pan and freeze until edges are solid.",
        "Transfer to processor and blend until smooth again.",
        "Fold in remaining 2 cups of blueberries.",
        "Pour into plastic mold and freeze overnight. Let soften slightly to serve."
    ],
    dietary: [ ]
} ]);
db.getCollection("recipe").insert([ {
    _id: ObjectId("6336eb430d300000ef083349"),
    "recipe_id": NumberLong("39"),
    name: "Biryani",
    "cook_time": NumberLong("1500"),
    "perp_time": NumberLong("14400"),
    "total_time": NumberLong("15900"),
    description: "Make and share this Biryani recipe from Food.com.",
    images: [
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/39/picM9Mhnw.jpg",
        "https://img.sndimg.com/food/image/upload/w_555,h_416,c_fit,fl_progressive,q_95/v1/img/recipes/39/picHv4Ocr.jpg"
    ],
    category: "Chicken Breast",
    keywords: [
        "Chicken Thigh & Leg",
        "Chicken",
        "Poultry",
        "Meat",
        "Asian",
        "Indian",
        "Weeknight",
        "Stove Top"
    ],
    ingredients: [
        "saffron",
        "milk",
        "hot green chili peppers",
        "onions",
        "garlic",
        "clove",
        "peppercorns",
        "cardamom seed",
        "cumin seed",
        "poppy seed",
        "mace",
        "cilantro",
        "mint leaf",
        "fresh lemon juice",
        "plain yogurt",
        "boneless chicken",
        "salt",
        "ghee",
        "onion",
        "tomatoes",
        "basmati rice",
        "long-grain rice",
        "raisins",
        "cashews",
        "eggs"
    ],
    calories: 1110.7,
    fat: 58.8,
    "saturated_fat": 16.6,
    cholesterol: 372.8,
    sodium: 368.4,
    carbonhydrate: 84.4,
    fiber: 9,
    sugar: 20.4,
    protein: 63.4,
    instruction: [
        "Soak saffron in warm milk for 5 minutes and puree in blender.",
        "Add chiles, onions, ginger, garlic, cloves, peppercorns, cardamom seeds, cinnamon, coriander and cumin seeds, poppy seeds, nutmeg, mace, cilantro or mint leaves and lemon juice. Blend into smooth paste. Put paste into large bowl, add yogurt and mix well.",
        "Marinate chicken in yogurt mixture with salt, covered for at least 2 - 6 hours in refrigerator.",
        "In skillet. heat oil over medium heat for 1 minute. Add ghee and 15 seconds later add onion and fry for about8 minutes.",
        "Reserve for garnish.",
        "In same skillet, cook chicken with its marinade with tomatoes for about 10 minutes over medium heat, uncovered.",
        "Remove chicken pieces from the sauce and set aside. Add rice to sauce, bring to boil, and cook, covered over low heat for 15 minutes.",
        "Return chicken and add raisins, cashews and almonds; mix well.",
        "Simmer, covered for 5 minutes.",
        "Place chicken, eggs and rice in large serving dish in such a way that yellow of the eggs, the saffron-colored rice, the nuts and the chicken make a colorful display.",
        "Add reserved onion as garnish."
    ],
    dietary: [
        "non-vegan",
        "non-vegetarian"
    ]
} ]);