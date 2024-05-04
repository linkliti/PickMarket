import PrioritySelector from "@/components/filters/PrioritySelector";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Filter } from "@/proto/app/protos/items";
import { useFiltersStore } from "@/store/filtersStore";
import { FilterPrefValue, FiltersStore } from "@/types/filterTypes";
import { ReactElement } from "react";
import terminal from "virtual:terminal";

export default function FilterSelectorBool({ filter }: { filter: Filter }): ReactElement {
  const [prefs, modifyPref] = useFiltersStore((state: FiltersStore) => [
    state.prefs,
    state.modifyPref,
  ]);

  function updateBoolSelector(): void {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "listVal") {
      const newValue: string[] = v.listVal.values.includes("Да") ? ["Нет"] : ["Да"];
      modifyPref(filter.key, {
        priority: prefs[filter.key].priority,
        value: {
          oneofKind: "listVal",
          listVal: {
            values: newValue,
          },
        },
      });
      terminal.log("Updated filter", filter.key, newValue);
    }
  }

  function isChecked(): boolean | undefined {
    const v: FilterPrefValue = prefs[filter.key].value;
    if (v.oneofKind === "listVal") {
      return v.listVal.values.includes("Да");
    }
    return false;
  }

  return (
    <Button
      asChild
      className="p-0"
    >
      <div className="bg-secondary flex w-full items-center justify-between overflow-hidden">
        <Label
          htmlFor={filter.key}
          className="w-10/12 truncate p-4 text-left"
        >
          {filter.title}
        </Label>
        <PrioritySelector key={filter.key} />
        <Switch
          className="mx-2 border border-gray-400"
          id={filter.key}
          checked={isChecked()}
          onCheckedChange={updateBoolSelector}
        />
      </div>
    </Button>
  );
}
