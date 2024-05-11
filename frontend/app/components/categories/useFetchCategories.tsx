import { Category } from "@/proto/app/protos/categories";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function useFetchCategories(marketLabel: string) {
  const [categories, setCategories] = useState<Category[]>([]);
  const [categoryURL, setCategoryURL] = useState<string>(`/api/categories/${marketLabel}/root`);

  const {
    isPending: isLoading,
    error,
    data: recentCategories,
  } = useQuery({
    queryKey: ["categories", categoryURL],
    queryFn: getCategories,
    staleTime: Infinity,
  });

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

  useEffect((): void => {
    setCategories([]);
  }, [marketLabel]);

  useEffect((): void => {
    if (!recentCategories) {
      return;
    }
    recentCategories.sort((a: Category, b: Category): number => a.title.localeCompare(b.title));
    // if (selectedCategory) {
    //   selectedCategory.isParsed = true;
    // }
    setCategories((prevCategories: Category[]): Category[] => [
      ...new Map(
        [...prevCategories, ...recentCategories].map((item: Category): [string, Category] => [
          item.url,
          item,
        ]),
      ).values(),
    ]);
  }, [recentCategories]);

  return {
    categories,
    isLoading,
    error,
    setCategoryURL,
  };
}
