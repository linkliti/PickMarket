import CategoryItem from "@/components/categories/CategoryItem";
import findParents from "@/components/categories/findParents";
import useFetchCategories from "@/components/categories/useFetchCategories";
import { RadioGroup } from "@/components/ui/radio-group";
import { Category } from "@/proto/app/protos/categories";
import { CategoryStore, useCategoryStore } from "@/store/categoryStore";
import { Marketplace } from "@/types/categoryTypes";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { ReactElement, useEffect } from "react";
import terminal from "virtual:terminal";

export default function CategorySelect({
  marketplace,
}: {
  marketplace: Marketplace;
}): ReactElement {
  const { categories, isLoading, error, setCategoryURL } = useFetchCategories(
    marketplace.shortLabel,
  );

  const [selectedCategory, setSelectedCategory] = useCategoryStore((state: CategoryStore) => [
    state.selectedCategory,
    state.setSelectedCategory,
  ]);

  useEffect((): void => {
    setSelectedCategory(null);
  }, [marketplace.shortLabel, setSelectedCategory]);

  function handleCategoryChange(category: Category): void {
    setCategoryURL(
      `/api/categories/${marketplace.shortLabel}/sub?url=${encodeURIComponent(category.url)}`,
    );
    setSelectedCategory(category);
    terminal.log("Selected", category.title);
  }

  if (isLoading && !categories) {
    return (
      <div className="flex items-center gap-2">
        <LoadingSpinner /> <p>Загрузка категорий</p>
      </div>
    );
  }
  if (error) {
    return (
      <div className="flex items-center gap-2">
        <p>Ошибка при загрузке категорий: {error?.message}</p>
      </div>
    );
  }
  return (
    <>
      <p className="mb-4">
        {"Выбранная категория: "}
        {selectedCategory?.title ? selectedCategory.title : "Не выбрана"}
      </p>
      <RadioGroup>{CategoryDisplay(categories)}</RadioGroup>
    </>
  );

  function renderCategoryItem(category: Category, level: number): ReactElement {
    return (
      <CategoryItem
        key={category.url}
        category={category}
        level={level}
        handleCategoryChange={handleCategoryChange}
        selectedCategory={selectedCategory}
      />
    );
  }

  function CategoryDisplay(categories: Category[]): ReactElement {
    if (categories.length === 0) {
      return <p>Не удалось загрузить категории</p>;
    }
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
                key={category.url}
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

    // Chain of parents
    let parents: Category[] = [];
    if (selectedCategory.parentUrl) {
      parents = findParents(selectedCategory.parentUrl, categories);
    }
    // Neighbours arrays
    const neighbours: Category[] = categories.filter(
      (category: Category): boolean => category.parentUrl === selectedCategory?.parentUrl,
    );
    const startNeighbours: Category[] = neighbours.slice(
      0,
      neighbours.findIndex((category: Category): boolean => category.url === selectedCategory?.url),
    );
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
        {parents.map((category: Category): ReactElement => renderCategoryItem(category, level++))}
        {startNeighbours.map(
          (category: Category): ReactElement => renderCategoryItem(category, level),
        )}
        {renderCategoryItem(selectedCategory, level)}
        {children.length > 0 &&
          children.map(
            (category: Category): ReactElement => renderCategoryItem(category, level + 1),
          )}
        {endNeighbours.map(
          (category: Category): ReactElement => renderCategoryItem(category, level),
        )}
      </>
    );
  }
}
