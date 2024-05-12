import { ItemsRequestWithPrefs } from "@/proto/app/protos/reqHandlerTypes";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { PrefForm, PrefFormSchema } from "@/types/filterTypes";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import terminal from "virtual:terminal";

export default function useFilterForm() {
  const [isOpen, setIsOpen] = useState<boolean>(true);
  const [formPrefs, setActivePrefs, setFormPrefs, market, categoryUrl] = useFilterStore(
    (state: FilterStore) => [
      state.formPrefs,
      state.setActivePrefs,
      state.setFormPrefs,
      state.market,
      state.categoryUrl,
    ],
  );
  const {
    handleSubmit,
    control,
    reset,
    formState: { isSubmitting },
  } = useForm<PrefForm>({
    defaultValues: formPrefs ? formPrefs : {},
    resolver: zodResolver(PrefFormSchema),
  });

  function onSubmit(data: PrefForm): void {
    terminal.log("Submitting form");
    const req: ItemsRequestWithPrefs = {
      request: {
        market: market,
        pageUrl: categoryUrl,
        numOfPages: data.numOfPages,
        params: "",
        userQuery: data.userQuery || "",
      },
      prefs: {},
    };

    for (const [key, priority] of Object.entries(data.priorities)) {
      if (!priority) continue;
      const charData: number | boolean | string[] = data.prefs[key];
      switch (typeof charData) {
        case "number": {
          req.prefs[key] = { priority: priority, value: { oneofKind: "numVal", numVal: charData } };
          break;
        }
        case "boolean": {
          req.prefs[key] = {
            priority: priority,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: charData ? ["Да"] : ["Нет"],
              },
            },
          };
          break;
        }
        case "object": {
          if (!charData.length) break;
          req.prefs[key] = {
            priority: priority,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: charData,
              },
            },
          };
        }
      }
    }
    setFormPrefs(data);
    setActivePrefs(req);
    setIsOpen(false);
  }

  return { isOpen, handleSubmit, control, reset, onSubmit, isSubmitting, setIsOpen, setFormPrefs };
}
