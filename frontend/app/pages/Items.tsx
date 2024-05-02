import FiltersNotSelected from "@/components/filters/FiltersNotSelected";
import FiltersSection from "@/components/filters/FiltersSection";
import { ReactElement } from "react";
import { useSearchParams } from "react-router-dom";

export default function ItemsPage(): ReactElement {
  const [searchParams] = useSearchParams();

  const market = searchParams.get("market");
  const categoryURL = searchParams.get("category");
  if (!market || !categoryURL) {
    return <FiltersNotSelected />;
  }

  return (
    <FiltersSection
      market={market}
      category={categoryURL}
    />
  );
}
