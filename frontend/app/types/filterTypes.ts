import { z } from "zod";

export const PrefFormSchema = z.object({
  params: z.string().optional(),
  userQuery: z.string().optional(),
  numOfPages: z
    .number({message: "Не задано количество страниц"})
    .min(1, "Количество страниц должно быть больше 0")
    .max(5, "Количество страниц не может быть больше 5"),
  priorities: z
    .record(z.number())
    .refine(AtleastOneValue, "Не настроены параметры приоритетов (звездочки)"),
  prefs: z
    .record(z.union([z.number(), z.array(z.string()), z.boolean()]))
    .refine(AtleastOneKey, "Не настроены параметры предпочтений"),
});

function AtleastOneKey(data: Record<string, unknown>): boolean {
  return Object.keys(data).length > 1;
}

function AtleastOneValue(data: Record<string, unknown>): boolean {
  return Object.values(data).some((value) => value);
}

export type PrefForm = z.infer<typeof PrefFormSchema>;
