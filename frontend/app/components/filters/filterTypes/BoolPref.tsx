import PrioritySelector from "@/components/filters/filterTypes/PrioritySelector";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { PrefForm } from "@/types/filterTypes";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function BoolPref({
  control,
  filterTitle,
  keyName,
}: {
  control: Control<PrefForm, unknown>;
  filterTitle: string;
  keyName: string;
}): ReactElement {
  return (
    <Controller
      name={`prefs.${keyName}`}
      control={control}
      defaultValue={false}
      render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => {
        const transformedValue: boolean | undefined =
          typeof value === "boolean" ? value : undefined;
        return (
          <Label
            htmlFor={keyName}
            className="ring-offset-background focus-visible:ring-ring text-primary-foreground hover:bg-primary/70 bg-secondary flex h-10 w-full cursor-pointer items-center justify-between gap-2 overflow-hidden whitespace-nowrap rounded-md px-4 py-2 text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 "
          >
            <p className="flex grow">{filterTitle}</p>
            <Switch
              className="border border-gray-400"
              id={keyName}
              name={name}
              ref={ref}
              checked={transformedValue}
              onCheckedChange={onChange}
              onBlur={onBlur}
              disabled={disabled}
            />
            <PrioritySelector
              key={keyName}
              control={control}
              keyName={keyName}
            />
          </Label>
        );
      }}
    />
  );
}
