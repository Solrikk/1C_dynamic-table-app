# Product Import Tool from Website to CSV

## Overview
This project is a Go language tool designed to automatically download, process, and save product data from a remote server into a CSV file. It is built for handling XML files containing information about products and their categories fetched from an e-commerce store. By converting data from XML to an easily readable and analyzable CSV format, this tool can be extremely useful for store owners, marketers, and analysts for further data processing and analysis.

## Features
- Downloads product data in XML format from a specified URL.
- Parses XML to extract detailed information about each product, including ID, name, category, price, old price, image URLs, and sizes.
- Saves the extracted information into a CSV file with a defined structure.
- Automatically creates a directory for saving the output if it does not exist.
