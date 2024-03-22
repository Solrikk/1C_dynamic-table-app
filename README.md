# BlueMoon: A Product Import Tool from Website to CSV

## Overview
The BlueMoon project is a tool developed in Go, aimed at automatically downloading, processing, and saving product data from a remote server into a CSV file format. Designed to handle XML files containing product and category information from an e-commerce platform, BlueMoon offers a seamless solution for transforming XML data into a CSV format that's easier to read and analyze. This makes it an invaluable resource for e-commerce store owners, marketing professionals, and data analysts for further data manipulation and analysis tasks.

## Features
- **Data Download**: Downloads product data in XML format from a provided URL.
- **XML Parsing**: Efficiently parses XML to extract comprehensive details on each product, including identifiers, names, categories, pricing information, URLs for images, and product sizes.
- **CSV Conversion**: Converts and saves the extracted product information into a neatly structured CSV file.
- **Directory Management**: Automatically ensures the existence of or creates a new directory for storing the exported CSV file, facilitating a hassle-free data export experience.
