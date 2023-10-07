from aocd import get_data

data = get_data(day=21, year=2020).splitlines()
allergen_ingredients = {}
foods = []

for line in data:
    ingredients, allergens = line.split(" (contains ")
    allergens = allergens[:-1].split(", ")
    foods.append([allergens, ingredients.split()])
# foods = [[['eggs', 'shellfish', ...], ['fljn', 'jcks', 'gxmptfk', 'tpjmxd', 'qdlpbt']], ...]

for allergen in {allergen for element in foods for allergen in element[0]}:
    common = []
    for food in foods:
        if allergen in food[0]:
            common = (
                food[1] if common == [] else list(set(common).intersection(food[1]))
            )
    allergen_ingredients[allergen] = common
# allergen_ingredients = {'shellfish': ['qdlpbt', 'pmvfzk', 'tsnkknk', 'kqv'], ... }

change_made = True
while change_made:
    change_made = False
    for allergen1 in allergen_ingredients:
        if len(allergen_ingredients[allergen1]) == 1:
            for allergen2 in allergen_ingredients:
                if (
                    allergen1 != allergen2
                    and allergen_ingredients[allergen1][0]
                    in allergen_ingredients[allergen2]
                ):
                    allergen_ingredients[allergen2].remove(
                        allergen_ingredients[allergen1][0]
                    )
                    change_made = True
# allergen_ingredients = {'shellfish': ['qdlpbt'], ... }


print("2020 Day 21")
print(
    "\tPart 1:",
    sum(
        [ingredient] not in allergen_ingredients.values()
        for ingredient in [ingredient for element in foods for ingredient in element[1]]
    ),
)
print(
    "\tPart 2:",
    ",".join(
        allergen_ingredients[allergen][0]
        for allergen in sorted(allergen_ingredients.keys())
    ),
)
