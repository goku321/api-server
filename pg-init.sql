CREATE TABLE geolocations (
		ip character varying NOT NULL,
		country_code character varying,
		country character varying,
		city character varying,
		latitude double precision,
		longitude double precision,
		mystery_value bigint,
		CONSTRAINT geolocations_pkey PRIMARY KEY (ip)
);

INSERT INTO "public"."geolocations"("ip", "country_code", "country", "city", "latitude", "longitude", "mystery_value")
VALUES
('200.106.141.15','SI','Nepal','DuBuquemouth',-84.87503094689836,7.206435933364332,7823011346);