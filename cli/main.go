package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pair-project/config"
	"pair-project/entity"
	"pair-project/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Selamat datang di Cakra Store")
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Masukkan pilihan (1/2/3): ")

		var option int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
		if err != nil {
			fmt.Println("Input anda bukan merupakan integer")
			continue
		}

		switch option {
		case 1:
			var email, password, firstName, lastName string
			fmt.Printf("\nREGISTER\n")
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&email)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			fmt.Print("Masukkan nama depan: ")
			fmt.Scanln(&firstName)
			fmt.Print("Masukkan nama belakang: ")
			fmt.Scanln(&lastName)

			user := entity.User{
				Email:     email,
				Password:  password,
				FirstName: firstName,
				LastName:  lastName,
			}

			if err := handler.Register(user, db); err != nil {
				fmt.Println("Kesalahan saat registrasi", err)
			} else {
				fmt.Println("Registrasi berhasil!")
			}

		case 2:
			var email, password string
			fmt.Printf("\nLOGIN\n")
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			user, authenticated, err := handler.Login(email, password, db)
			if err != nil {
				panic(err.Error())
			} else if authenticated {
				for {
					fmt.Printf("\nSelamat datang %v. Pilih opsi berikut:\n", user.Email)
					fmt.Println("1. Tampilkan semua produk")
					fmt.Println("2. Order pesanan")
					fmt.Println("3. Tampilkan riwayat pesanan")
					fmt.Println("4. Hapus riwayat pesanan")
					fmt.Println("5. View User Report")
					fmt.Println("6. View Order Report")
					fmt.Println("7. View Stock Report")
					fmt.Println("8. Exit")
					fmt.Print("Masukkan pilihan (1/2/3/4/5/6/7/8): ")

					var option int
					scanner.Scan()
					_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
					if err != nil {
						fmt.Println("Input anda bukan merupakan integer")
						continue
					}

					switch option {
					case 1:
						err := handler.DisplayProducts(db)
						if err != nil {
							log.Fatal(err)
						}

					case 2:
						var productOption, quantity int
						fmt.Print("Masukkan ID product: ")
						scanner.Scan()
						_, err := fmt.Sscanf(scanner.Text(), "%d", &productOption)
						if err != nil {
							log.Fatal("Input bukan merupakan angka")
						}

						fmt.Print("Masukkan Quantity: ")
						scanner.Scan()
						_, err = fmt.Sscanf(scanner.Text(), "%d", &quantity)
						if err != nil {
							log.Fatal("Input bukan merupakan angka")
						}

						err = handler.PlaceOrder(db, user, productOption, quantity)
						if err != nil {
							log.Fatal(err)
						}

					case 3:
						err := handler.DisplayOrderHistory(db, user)
						if err != nil {
							log.Fatal(err)
						}

					case 4:
						fmt.Print("Masukkan Order ID yang ingin dihapus: ")
						var orderID int
						scanner.Scan()
						_, err := fmt.Sscanf(scanner.Text(), "%d", &orderID)
						if err != nil {
							log.Fatal("Input bukan merupakan angka")
						}

						err = handler.DeleteOrderHistory(db, user, orderID)
						if err != nil {
							log.Fatal(err)
						}

					case 5:
						err := handler.ViewUserReport(db, user.Id)
						if err != nil {
							log.Fatal(err)
						}

					case 6:
						err := handler.ViewOrderReport(db, user.Id)
						if err != nil {
							log.Fatal(err)
						}

					case 7:
						err := handler.ViewStockReport(db)
						if err != nil {
							log.Fatal(err)
						}

					case 8:
						fmt.Println("Sampai jumpa!")
						os.Exit(0)

					default:
						fmt.Println("Input harus merupakan angka 1-5")
						os.Exit(1)
					}
				}
			} else {
				fmt.Println("\nFailed to login")
				fmt.Println("=============================")
			}

		case 3:
			fmt.Println("Sampai jumpa!")
			os.Exit(0)

		default:
			fmt.Println("Input harus merupakan angka 1-3!")
		}
	}
}
