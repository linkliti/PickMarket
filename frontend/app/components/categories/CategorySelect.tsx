import CategoryItem from "@/components/categories/CategoryItem";
import { RadioGroup } from "@/components/ui/radio-group";
import { Category } from "@/types/categoryTypes";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function CategorySelect(): ReactElement {
  const [categories, setCategories] = useState<Category[]>([]);
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null);

  useEffect((): void => {
    fetchCategories("/api/categories/ozon/root");
  }, []);

  function fetchCategories(url: string): void {
    fetch(url)
      .then((res: Response): Promise<Category[]> => res.json())
      .then((data: Category[]): void => {
        if (data) {
          data.sort((a: Category, b: Category): number => a.title.localeCompare(b.title));
          setCategories((prevCategories: Category[]): Category[] => [
            ...new Map(
              [...prevCategories, ...data].map((item: Category): [string, Category] => [
                item.url,
                item,
              ]),
            ).values(),
          ]);
        }
      })
      .catch(terminal.error);
  }

  function handleCategoryChange(category: Category): void {
    if (!category.isParsed) {
      terminal.log("Fetching", category.url);
      fetchCategories(`/api/categories/ozon/sub?url=${encodeURIComponent(category.url)}`);
      category.isParsed = true;
    }
    setSelectedCategory(category);
    terminal.log("Selected", category.title);
  }

  function displayCategories(categories: Category[]): ReactElement {
    let level: number = 0;
    // Если не выбрана категория, отображаем только категории без родителей
    if (!selectedCategory) {
      const filteredCategories: Category[] = categories.filter(
        (category: Category): boolean => !category.parentUrl,
      );
      return (
        <>
          {filteredCategories.map(
            (category: Category): ReactElement => (
              <CategoryItem
                category={category}
                level={level}
                handleCategoryChange={handleCategoryChange}
              />
            ),
          )}
        </>
      );
    }

    // Parents array of selectedCategory for nesting
    function findParents(parentUrl: string): Category[] {
      const parentCategories: Category[] = [];
      for (;;) {
        const parentCat: Category = categories.filter(
          (category: Category): boolean => parentUrl === category.url,
        )[0];
        parentCategories.unshift(parentCat);
        if (parentCat.parentUrl) {
          parentUrl = parentCat.parentUrl;
        } else {
          break;
        }
      }
      return parentCategories;
    }

    // Chain of parents
    let parents: Category[] = [];
    if (selectedCategory.parentUrl) {
      parents = findParents(selectedCategory.parentUrl);
    }
    // Neighbours arrays
    const neighbours: Category[] = categories.filter(
      (category: Category): boolean => category.parentUrl === selectedCategory?.parentUrl,
    );
    const startNeighbours: Category[] = neighbours.slice(
      0,
      neighbours.findIndex((category: Category): boolean => category.url === selectedCategory?.url),
    );
    const selectedCat: Category = categories.filter(
      (category: Category): boolean => category.url === selectedCategory?.url,
    )[0];
    const endNeighbours: Category[] = neighbours.slice(
      neighbours.findIndex(
        (category: Category): boolean => category.url === selectedCategory?.url,
      ) + 1,
    );

    // Children array
    const children: Category[] = categories.filter(
      (category: Category): boolean => category.parentUrl === selectedCategory?.url,
    );

    // Return
    return (
      <>
        {parents.map(
          (category: Category): ReactElement => (
            <CategoryItem
              category={category}
              level={level++}
              handleCategoryChange={handleCategoryChange}
              selectedCategory={selectedCategory}
            />
          ),
        )}
        {startNeighbours.map(
          (category: Category): ReactElement => (
            <CategoryItem
              category={category}
              level={level}
              handleCategoryChange={handleCategoryChange}
              selectedCategory={selectedCategory}
            />
          ),
        )}
        <CategoryItem
          category={selectedCat}
          level={level}
          handleCategoryChange={handleCategoryChange}
          selectedCategory={selectedCategory}
        />
        {children.length > 0 &&
          children.map(
            (category: Category): ReactElement => (
              <CategoryItem
                category={category}
                level={level + 1}
                handleCategoryChange={handleCategoryChange}
                selectedCategory={selectedCategory}
              />
            ),
          )}
        {endNeighbours.map(
          (category: Category): ReactElement => (
            <CategoryItem
              category={category}
              level={level}
              handleCategoryChange={handleCategoryChange}
              selectedCategory={selectedCategory}
            />
          ),
        )}
      </>
    );
  }

  return (
    <>
      <p className="mb-4">
        Выбранная категория: {selectedCategory?.title ? selectedCategory.title : "Не выбрана"}
      </p>
      <RadioGroup>{displayCategories(categories)}</RadioGroup>
    </>
  );
}
