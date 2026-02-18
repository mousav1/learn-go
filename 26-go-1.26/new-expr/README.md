# Go 1.26 — new(expr)

در Go 1.26 می‌توانید expression را مستقیماً به new بدهید.

قبلاً:

v := 10
p := &v

حالا:

p := new(10)

مزایا:

- کد کوتاه‌تر
- خواناتر
- بدون متغیر موقت
- مناسب برای JSON / struct pointer fields
