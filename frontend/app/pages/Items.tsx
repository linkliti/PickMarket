import FiltersNotSelected from "@/components/filters/FiltersNotSelected";
import FiltersSection from "@/components/filters/FiltersSection";
import ItemsSection from "@/components/items/ItemsSection";
import { Markets } from "@/proto/app/protos/types";
import { useFilterStore } from "@/store/filterStore";
import { ReactElement, useEffect } from "react";
import { useSearchParams } from "react-router-dom";

export default function ItemsPage(): ReactElement {
  useEffect((): void => {
    document.title = "Предпочтения";
  }, []);

  const [searchParams] = useSearchParams();
  const [setMarket, setCategoryUrl] = useFilterStore((state) => [
    state.setMarket,
    state.setCategoryUrl,
  ]);

  const marketStr: string | null = searchParams.get("market");
  const categoryURL: string | null = searchParams.get("category");

  useEffect(() => {
    if (!marketStr || !categoryURL) {
      return;
    }
    const market: number = Number(marketStr);
    setMarket(market);
    setCategoryUrl(categoryURL);
  }, [marketStr, categoryURL, setMarket, setCategoryUrl]);

  if (!marketStr || !categoryURL || !Markets[Number(marketStr)]) {
    return <FiltersNotSelected />;
  }
  return (
    <>
      <FiltersSection />
      <ItemsSection />
    </>
  );
}
