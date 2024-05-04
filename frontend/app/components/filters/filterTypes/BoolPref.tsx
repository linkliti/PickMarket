import FilterWrapper from "@/components/filters/filterTypes/FilterWrapper";
import { Switch } from "@/components/ui/switch";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function BoolPref({
  control,
  name,
  keyName,
}: {
  control: Control;
  name: string;
  keyName: string;
}): ReactElement {
  return (
    <FilterWrapper
      name={name}
      keyName={keyName}
    >
      <Controller
        name={keyName}
        control={control}
        defaultValue={false}
        render={({ field }) => <Switch {...field} />}
      />
    </FilterWrapper>
  );
}
