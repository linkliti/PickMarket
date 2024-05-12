import { FilterContext } from "@/components/items/FilterContext";
import { ItemContext } from "@/components/items/ItemContext";
import { Button } from "@/components/ui/button";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import filtersLink from "@/utilities/toQuery";
import { ReactElement, useContext } from "react";
import { useNavigate } from "react-router-dom";
import terminal from "virtual:terminal";

export default function ItemGoToSavedFilters(): ReactElement {
  const [setFormPrefs] = useFilterStore((state: FilterStore) => [state.setFormPrefs]);
  const { categoryURL, form } = useContext(FilterContext);
  const { market } = useContext(ItemContext);

  const navigate = useNavigate();

  function goToSavedFilters() {
    const url = filtersLink(market, categoryURL);
    terminal.log("Navigating to " + url);
    setFormPrefs(form);
    navigate(url);
  }

  return (
    <Button
      className="mt-6 text-wrap py-6 max-md:mt-2"
      onClick={goToSavedFilters}
    >
      Перейти к сохраненным предпочтениям
    </Button>
  );
}
