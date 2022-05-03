


Suppose an e-commerce site wants to serve customers in the US, EU, and Asia. However, EU regulations require GDPR compliance
that do not apply to the US & Asia. Chine, however, has strict data residency requirements that do not apply to any other country, but must
be respected to enter the Chinese market. With the GDN, a dedicated EU data fabric allows the company to comply with EU regulations while
storing and processing data of EU customers only on EU territory. Likewise, a dedicated US data fabric ensures low latency especially when
combined with Cloudflare cache while ascertain that US data & privacy regulations only apply to US customers. For Asia,
a third data fabric may cover all of South East Asia except China.

An internal data analytic service then pulls data from the US & Asia fabric, but may not access the EU data fabric
unless EU private data remain masked at all times. For performance reasons, the data analytics cluster resides in a POP
shared by both the US & Asia data fabric as to stream real-time analytics back to the eCommerce store.

Because of strict data residence regulations in China, a private cloud deployment on a Chinese cloud provider
would then separate data from chinese users from the rest of the world while complying with Chinese regulations.
Note, in any case, seek legal advice on national and international privacy & data processing before entering
regulated markets. 