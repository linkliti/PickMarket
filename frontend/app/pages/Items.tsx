import FiltersNotSelected from "@/components/filters/FiltersNotSelected";
import FiltersSection from "@/components/filters/FiltersSection";
import { Markets } from "@/proto/app/protos/types";
import { ReactElement } from "react";
import { useSearchParams } from "react-router-dom";
import terminal from "virtual:terminal";

export default function ItemsPage(): ReactElement {
  const [searchParams] = useSearchParams();

  const marketStr: string | null = searchParams.get("market");
  const categoryURL: string | null = searchParams.get("category");
  terminal.log(marketStr, categoryURL);
  if (!marketStr || !categoryURL) {
    return <FiltersNotSelected />;
  }

  const market: number = Number(marketStr);

  if (!Markets[market]) {
    return <FiltersNotSelected />;
  }
  return (
    <FiltersSection
      market={market}
      category={categoryURL}
    />
  );
}
