package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/QOM-One/QomApp/v7/app"
	"github.com/QOM-One/QomApp/v7/contracts"
	"github.com/QOM-One/QomApp/v7/x/erc20/types"
	onboardingtest "github.com/QOM-One/QomApp/v7/x/onboarding/testutil"
)

var _ = Describe("Onboarding: Performing an IBC Transfer followed by autoswap and convert", Ordered, func() {
	coincanto := sdk.NewCoin("aqom", sdk.ZeroInt())
	ibcBalance := sdk.NewCoin(uusdcIbcdenom, sdk.NewIntWithDecimal(10000, 6))
	coinUsdc := sdk.NewCoin("uUSDC", sdk.NewIntWithDecimal(10000, 6))
	coinAtom := sdk.NewCoin("uatom", sdk.NewIntWithDecimal(10000, 6))

	var (
		sender, receiver string
		senderAcc        sdk.AccAddress
		receiverAcc      sdk.AccAddress
		result           *sdk.Result
		tokenPair        *types.TokenPair
	)

	BeforeEach(func() {
		s.SetupTest()
	})

	Describe("from a non-authorized channel: Cosmos ---(uatom)---> Qom", func() {
		BeforeEach(func() {
			// deploy ERC20 contract and register token pair
			tokenPair = s.setupRegisterCoin(metadataIbcUSDC)

			// send coins from Cosmos to canto
			sender = s.IBCCosmosChain.SenderAccount.GetAddress().String()
			receiver = s.cantoChain.SenderAccount.GetAddress().String()
			senderAcc = sdk.MustAccAddressFromBech32(sender)
			receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			result = s.SendAndReceiveMessage(s.pathCosmosqom, s.IBCCosmosChain, "uatom", 10000000000, sender, receiver, 1)

		})
		It("No swap and convert operation - aqom balance should be 0", func() {
			nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
			Expect(nativecanto).To(Equal(coincanto))
		})
		It("Qom chain's IBC voucher balance should be same with the transferred amount", func() {
			ibcAtom := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uatomIbcdenom)
			Expect(ibcAtom).To(Equal(sdk.NewCoin(uatomIbcdenom, coinAtom.Amount)))
		})
		It("Cosmos chain's uatom balance should be 0", func() {
			atom := s.IBCCosmosChain.GetSimApp().BankKeeper.GetBalance(s.IBCCosmosChain.GetContext(), senderAcc, "uatom")
			Expect(atom).To(Equal(sdk.NewCoin("uatom", sdk.ZeroInt())))
		})
	})

	Describe("from an authorized channel: Gravity ---(uUSDC)---> Qom", func() {
		When("ERC20 contract is deployed and token pair is enabled", func() {
			BeforeEach(func() {
				// deploy ERC20 contract and register token pair
				tokenPair = s.setupRegisterCoin(metadataIbcUSDC)

				sender = s.IBCGravityChain.SenderAccount.GetAddress().String()
				receiver = s.cantoChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)

				s.FundCantoChain(sdk.NewCoins(ibcBalance))

			})

			Context("when no swap pool exists", func() {
				BeforeEach(func() {
					result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)
				})
				It("No swap: aqom balance should be 0", func() {
					nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
					Expect(nativecanto).To(Equal(coincanto))
				})
				It("Convert: Qom chain's IBC voucher balance should be same with the original balance", func() {
					ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
					Expect(ibcUsdc).To(Equal(ibcBalance))
				})
				It("Convert: ERC20 token balance should be same with the transferred IBC voucher amount", func() {
					erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
					Expect(erc20balance).To(Equal(coinUsdc.Amount.BigInt()))
				})
				It("Convert: ERC20 token balance should be same with the converted IBC voucher amount", func() {
					events := result.GetEvents()
					attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "convert_coin"))
					convertAmount, _ := sdk.NewIntFromString(attrs["amount"])
					erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
					Expect(erc20balance).To(Equal(convertAmount.BigInt()))
				})
			})

			Context("when swap pool exists", func() {
				BeforeEach(func() {
					s.CreatePool(uusdcIbcdenom)
				})
				When("aqom balance is 0 and not enough IBC token transferred to swap aqom", func() {
					BeforeEach(func() {
						result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 1000000, sender, receiver, 1)
					})
					It("No swap: Balance of aqom should be same with the original aqom balance (0)", func() {
						nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
						Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", sdk.ZeroInt())))
					})
					It("Convert: Qom chain's IBC voucher balance should be same with the original balance", func() {
						ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
						Expect(ibcUsdc).To(Equal(ibcBalance))
					})
					It("Convert: ERC20 token balance should be same with the transferred IBC voucher amount", func() {
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(sdk.NewIntWithDecimal(1, 6).BigInt()))
					})
					It("Convert: ERC20 token balance should be same with the converted IBC voucher amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "convert_coin"))
						convertAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(convertAmount.BigInt()))
					})
				})

				When("Qom chain's aqom balance is 0", func() {
					BeforeEach(func() {
						result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)
					})
					It("Swap: balance of aqom should be same with the auto swap threshold", func() {
						autoSwapThreshold := s.cantoChain.App.(*app.Qom).OnboardingKeeper.GetParams(s.cantoChain.GetContext()).AutoSwapThreshold
						nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
						Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", autoSwapThreshold)))
					})
					It("Convert: Qom chain's IBC voucher balance should be same with the original balance", func() {
						ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
						Expect(ibcUsdc).To(Equal(ibcBalance))
					})
					It("Convert: ERC20 token balance should be same with the difference between transferred IBC voucher amount and the swapped amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "swap"))
						swappedAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(coinUsdc.Amount.Sub(swappedAmount).BigInt()))
					})
					It("Convert: ERC20 token balance should be same with the converted IBC voucher amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "convert_coin"))
						convertAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(convertAmount.BigInt()))
					})
				})

				When("Qom chain's aqom balance is between 0 and auto swap threshold (3canto)", func() {
					BeforeEach(func() {
						s.FundCantoChain(sdk.NewCoins(sdk.NewCoin("aqom", sdk.NewIntWithDecimal(3, 18))))
						result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)
					})
					It("Auto swap operation: balance of aqom should be same with the sum of original aqom balance and auto swap threshold", func() {
						autoSwapThreshold := s.cantoChain.App.(*app.Qom).OnboardingKeeper.GetParams(s.cantoChain.GetContext()).AutoSwapThreshold
						nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
						Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", autoSwapThreshold.Add(sdk.NewIntWithDecimal(3, 18)))))
					})
					It("Convert: Qom chain's IBC voucher balance should be same with the original balance", func() {
						ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
						Expect(ibcUsdc).To(Equal(ibcBalance))
					})
					It("Convert: ERC20 token balance should be same with the difference between transferred IBC voucher amount and the swapped amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "swap"))
						swappedAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(coinUsdc.Amount.Sub(swappedAmount).BigInt()))
					})
					It("Convert: ERC20 token balance should be same with the converted IBC voucher amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "convert_coin"))
						convertAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(convertAmount.BigInt()))
					})
				})
				When("Qom chain's aqom balance is bigger than the auto swap threshold (4canto)", func() {
					BeforeEach(func() {
						s.FundCantoChain(sdk.NewCoins(sdk.NewCoin("aqom", sdk.NewIntWithDecimal(4, 18))))
						result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)
					})
					It("No swap: balance of aqom should be same with the original aqom balance (4canto)", func() {
						nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
						Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", sdk.NewIntWithDecimal(4, 18))))
					})
					It("Convert: Qom chain's IBC voucher balance should be same with the original balance", func() {
						ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
						Expect(ibcUsdc).To(Equal(ibcBalance))
					})
					It("Convert: ERC20 token balance should be same with the transferred IBC voucher amount", func() {
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(coinUsdc.Amount.BigInt()))
					})
					It("Convert: ERC20 token balance should be same with the converted IBC voucher amount", func() {
						events := result.GetEvents()
						attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "convert_coin"))
						convertAmount, _ := sdk.NewIntFromString(attrs["amount"])
						erc20balance := s.cantoChain.App.(*app.Qom).Erc20Keeper.BalanceOf(s.cantoChain.GetContext(), contracts.ERC20MinterBurnerDecimalsContract.ABI, tokenPair.GetERC20Contract(), common.BytesToAddress(receiverAcc.Bytes()))
						Expect(erc20balance).To(Equal(convertAmount.BigInt()))
					})
				})
			})
		})
		When("ERC20 contract is deployed and token pair is enabled", func() {
			BeforeEach(func() {
				// deploy ERC20 contract and register token pair
				tokenPair = s.setupRegisterCoin(metadataIbcUSDC)
				tokenPair.Enabled = false
				s.cantoChain.App.(*app.Qom).Erc20Keeper.SetTokenPair(s.cantoChain.GetContext(), *tokenPair)
				sender = s.IBCGravityChain.SenderAccount.GetAddress().String()
				receiver = s.cantoChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)

				s.CreatePool(uusdcIbcdenom)
				s.FundCantoChain(sdk.NewCoins(ibcBalance))
				s.FundCantoChain(sdk.NewCoins(sdk.NewCoin("aqom", sdk.NewIntWithDecimal(3, 18))))
				result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)

			})
			It("Auto swap operation: balance of aqom should be same with the sum of original aqom balance and auto swap threshold", func() {
				autoSwapThreshold := s.cantoChain.App.(*app.Qom).OnboardingKeeper.GetParams(s.cantoChain.GetContext()).AutoSwapThreshold
				nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
				Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", autoSwapThreshold.Add(sdk.NewIntWithDecimal(3, 18)))))
			})
			It("No convert: Qom chain's IBC voucher balance should be same with (original balance + transferred amount - swapped amount)", func() {
				events := result.GetEvents()
				attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "swap"))
				swappedAmount, _ := sdk.NewIntFromString(attrs["amount"])
				ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
				Expect(ibcUsdc.Amount).To(Equal(ibcBalance.Amount.Add(sdk.NewInt(10000000000)).Sub(swappedAmount)))
			})
		})
		When("ERC20 contract is not deployed", func() {
			BeforeEach(func() {
				sender = s.IBCGravityChain.SenderAccount.GetAddress().String()
				receiver = s.cantoChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)

				s.CreatePool(uusdcIbcdenom)
				s.FundCantoChain(sdk.NewCoins(ibcBalance))
				s.FundCantoChain(sdk.NewCoins(sdk.NewCoin("aqom", sdk.NewIntWithDecimal(3, 18))))
				result = s.SendAndReceiveMessage(s.pathGravityqom, s.IBCGravityChain, "uUSDC", 10000000000, sender, receiver, 1)
			})
			It("Auto swap operation: balance of aqom should be same with the sum of original aqom balance and auto swap threshold", func() {
				autoSwapThreshold := s.cantoChain.App.(*app.Qom).OnboardingKeeper.GetParams(s.cantoChain.GetContext()).AutoSwapThreshold
				nativecanto := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, "aqom")
				Expect(nativecanto).To(Equal(sdk.NewCoin("aqom", autoSwapThreshold.Add(sdk.NewIntWithDecimal(3, 18)))))
			})
			It("No convert: Qom chain's IBC voucher balance should be same with (original balance + transferred amount - swapped amount)", func() {
				events := result.GetEvents()
				attrs := onboardingtest.ExtractAttributes(onboardingtest.FindEvent(events, "swap"))
				swappedAmount, _ := sdk.NewIntFromString(attrs["amount"])
				ibcUsdc := s.cantoChain.App.(*app.Qom).BankKeeper.GetBalance(s.cantoChain.GetContext(), receiverAcc, uusdcIbcdenom)
				Expect(ibcUsdc.Amount).To(Equal(ibcBalance.Amount.Add(sdk.NewInt(10000000000)).Sub(swappedAmount)))
			})

		})
	})

})
