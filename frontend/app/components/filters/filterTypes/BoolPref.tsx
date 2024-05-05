import PrioritySelector from "@/components/filters/filterTypes/PrioritySelector";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { cn } from "@/lib/utils";
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
          <Button
            asChild
            className={cn(
              "bg-secondary flex w-full items-center justify-between gap-2 overflow-hidden",
            )}
          >
            <div>
              <div className="grow truncate text-left">{filterTitle}</div>
              <PrioritySelector
                key={keyName}
                control={control}
                keyName={keyName}
              />
              <Switch
                className="border border-gray-400"
                name={name}
                ref={ref}
                checked={transformedValue}
                onCheckedChange={onChange}
                onBlur={onBlur}
                disabled={disabled}
              />
            </div>
          </Button>
        );
      }}
    />
  );
}
