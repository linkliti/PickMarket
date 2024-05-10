import CategoryItem from "@/components/categories/CategoryItem";
import { RadioGroup } from "@/components/ui/radio-group";
import { Category } from "@/proto/app/protos/categories";
import { CategoryStore, useCategoryStore } from "@/store/categoryStore";
import { Marketplace } from "@/types/categoryTypes";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function CategorySelect({
  marketplace,
}: {
  marketplace: Marketplace;
}): ReactElement {
  const [categoryURL, setCategoryURL] = useState<string>(
    `/api/categories/${marketplace.shortLabel}/root`,
  );
  const [categories, setCategories] = useState<Category[]>([]);
  const [selectedCategory, setSelectedCategory] = useCategoryStore((state: CategoryStore) => [
    state.selectedCategory,
    state.setSelectedCategory,
  ]);

  useEffect((): void => {
    setCategories([]);
    setSelectedCategory(null);
  }, [marketplace.shortLabel, setSelectedCategory]);

  const {
    isPending: isLoading,
    error,
    data: recentCategories,
  } = useQuery({
    queryKey: ["categories", categoryURL],
    queryFn: getCategories,
    staleTime: Infinity,
  });

  useEffect((): void => {
    if (!recentCategories) {
      return;
    }
    recentCategories.sort((a: Category, b: Category): number => a.title.localeCompare(b.title));
    if (selectedCategory) {
      selectedCategory.isParsed = true;
    }
    setCategories((prevCategories: Category[]): Category[] => [
      ...new Map(
        [...prevCategories, ...recentCategories].map((item: Category): [string, Category] => [
          item.url,
          item,
        ]),
      ).values(),
    ]);
  }, [recentCategories, selectedCategory]);

  async function getCategories(): Promise<Category[] | undefined> {
    try {
      terminal.log("Fetching", categoryURL);
      const res: AxiosResponse = await axios.get<Category[]>(categoryURL);
      const data: Category[] = res.data;
      return data;
    } catch (error) {
      if (error instanceof Error) {
        terminal.error(error.message);
        throw new Error(error.message);
      }
    }
  }

  function handleCategoryChange(category: Category): void {
    setCategoryURL(
      `/api/categories/${marketplace.shortLabel}/sub?url=${encodeURIComponent(category.url)}`,
    );
    setSelectedCategory(category);
    terminal.log("Selected", category.title);
  }

  function displayCategories(categories: Category[]): ReactElement {
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
              key={category.url}
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
              key={category.url}
              category={category}
              level={level}
              handleCategoryChange={handleCategoryChange}
              selectedCategory={selectedCategory}
            />
          ),
        )}
        <CategoryItem
          key={selectedCategory.url}
          category={selectedCat}
          level={level}
          handleCategoryChange={handleCategoryChange}
          selectedCategory={selectedCategory}
        />
        {children.length > 0 &&
          children.map(
            (category: Category): ReactElement => (
              <CategoryItem
                key={category.url}
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

  return (
    <>
      {isLoading && !categories ? (
        <div className="flex items-center gap-2">
          <LoadingSpinner /> <p>Загрузка категорий</p>
        </div>
      ) : error ? (
        <div className="flex items-center gap-2">
          <p>Ошибка при загрузке категорий: {error?.message}</p>
        </div>
      ) : (
        <>
          <p className="mb-4">
            {"Выбранная категория: "}
            {selectedCategory?.title ? selectedCategory.title : "Не выбрана"}
          </p>
          <RadioGroup>{displayCategories(categories)}</RadioGroup>
        </>
      )}
    </>
  );
}
