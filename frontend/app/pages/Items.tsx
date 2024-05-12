import Loading from "@/components/base/Loading";
import FiltersNotSelected from "@/components/filters/FiltersNotSelected";
import FiltersSection from "@/components/filters/FiltersSection";
import ItemsSection from "@/components/items/ItemsSection";
import { Markets } from "@/proto/app/protos/types";
import { useFilterStore } from "@/store/filterStore";
import { ReactElement, useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import terminal from "virtual:terminal";

export default function ItemsPage(): ReactElement {
  useEffect((): void => {
    document.title = "Предпочтения";
  }, []);

  const [searchParams] = useSearchParams();
  const [setMarket, setCategoryUrl, savedCategoryUrl] = useFilterStore((state) => [
    state.setMarket,
    state.setCategoryUrl,
    state.categoryUrl,
  ]);

  const marketStr: string | null = searchParams.get("market");
  const categoryURL: string | null = searchParams.get("category");

  useEffect(() => {
    if (!marketStr || !categoryURL) {
      return;
    }
    terminal.log(marketStr, categoryURL);
    const market: number = Number(marketStr);
    setMarket(market);
    setCategoryUrl(categoryURL);
  }, [marketStr, categoryURL, setMarket, setCategoryUrl]);

  if (!marketStr || !categoryURL || !Markets[Number(marketStr)]) {
    return <FiltersNotSelected />;
  }
  if (!savedCategoryUrl) {
    return <Loading message="Загрузка предпочтений" />;
  }
  return (
    <>
      <FiltersSection />
      <ItemsSection />
    </>
  );
}
