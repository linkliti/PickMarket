import BoolPref from "@/components/filters/filterTypes/BoolPref";
import RangePref from "@/components/filters/filterTypes/RangePref";
import SelectionPref from "@/components/filters/filterTypes/SelectionPref";
import { Option } from "@/components/ui/multiple-selector";
import { Filter } from "@/proto/app/protos/items";
import { SelectionFilterItem } from "@/proto/app/protos/types";
import { ReactElement } from "react";
import { Control } from "react-hook-form";

export default function FilterForm({
  filter,
  control,
}: {
  filter: Filter;
  control: Control;
}): ReactElement {
  switch (filter.data.oneofKind) {
    case "rangeFilter": {
      return (
        <RangePref
          key={filter.key}
          keyName={filter.key}
          control={control}
          name={filter.title}
          range={filter.data.rangeFilter}
        />
      );
    }
    case "selectionFilter": {
      const listVals: Option[] = Array.from(
        new Set(
          filter.data.selectionFilter.items.map((item: SelectionFilterItem): string => item.text),
        ),
      ).map((value: string): Option => ({ value, label: value }));

      return (
        <SelectionPref
          key={filter.key}
          keyName={filter.key}
          control={control}
          name={filter.title}
          options={listVals}
        />
      );
    }
    case "boolFilter": {
      return (
        <BoolPref
          key={filter.key}
          keyName={filter.key}
          control={control}
          name={filter.title}
        />
      );
    }
    default: {
      return <></>;
    }
  }
}
