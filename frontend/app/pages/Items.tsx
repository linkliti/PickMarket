import Filters from "@/components/filters/Filters";
import FiltersNotSelected from "@/components/filters/FiltersNotSelected";
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
    <Filters
      market={market}
      category={categoryURL}
    />
  );
}
