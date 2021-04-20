package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"

	"github.com/quycao/xstats/pkg/stock"
	"github.com/quycao/xstats/pkg/util"
)

var inputs = []string{
	// HOSE
	"AAA", "AAM", "AAT", "ABS", "ABT", "ACB", "ACC", "ACL", "ADG", "ADS", "AGD", "AGG", "AGM", "AGR", "AHP", "ALP", "AMD", "ANV", "APC", "APG", "APH", "ASG", "ASM", "ASP", "AST", "ATG",
	"BAS", "BBC", "BCE", "BCG", "BCI", "BCM", "BFC", "BGM", "BHN", "BIC", "BID", "BKG", "BMC", "BMI", "BMP", "BRC", "BSI", "BTP", "BTT", "BVH", "BWE",
	"C32", "C47", "CAV", "CCI", "CCL", "CDC", "CEE", "CHP", "CIG", "CII", "CKG", "CLC", "CLG", "CLL", "CLP", "CLW", "CMG", "CMV", "CMX", "CNG", "COM", "CRC", "CRE", "CSG", "CSM", "CSV", "CTD", "CTF", "CTG", "CTI", "CTS", "CVT",
	"D2D", "DAG", "DAH", "DAT", "DBC", "DBD", "DBT", "DC4", "DCC", "DCL", "DCM", "DGC", "DGW", "DHA", "DHC", "DHG", "DHM", "DIG", "DLG", "DMC", "DPG", "DPM", "DPR", "DQC", "DRC", "DRH", "DRL", "DSN", "DTA", "DTL", "DTT", "DVD", "DVP", "DXG", "DXV",
	"EIB", "ELC", "EMC", "EVE", "EVG", "FBT", "FCM", "FCN", "FDC", "FIR", "FIT", "FLC", "FMC", "FPC", "FPT", "FRT", "FTM", "FTS",
	"GAB", "GAS", "GDT", "GEG", "GEX", "GIL", "GMC", "GMD", "GSP", "GTA", "GTN", "GVR",
	"HAG", "HAH", "HAI", "HAP", "HAR", "HAS", "HAX", "HBC", "HCD", "HCM", "HDB", "HDC", "HDG", "HHP", "HHS", "HID", "HII", "HMC", "HNG", "HOT", "HPG", "HPX", "HQC", "HRC", "HSG", "HSL", "HT1", "HT2", "HTI", "HTL", "HTN", "HTV", "HU1", "HU3", "HUB", "HVH", "HVN", "HVX",
	"IBC", "ICT", "IDI", "IJC", "ILB", "IMP", "ITA", "ITC", "ITD", "JVC", "KBC", "KDC", "KDH", "KHP", "KMR", "KOS", "KPF", "KSB",
	"L10", "LAF", "LBM", "LCG", "LCM", "LDG", "LEC", "LGC", "LGL", "LHG", "LIX", "LM8", "LPB", "LSS",
	"MBB", "MCG", "MCP", "MCV", "MDG", "MHC", "MIG", "MSB", "MSH", "MSN", "MWG", "NAF", "NAV", "NBB", "NCT", "NHA", "NHH", "NHS", "NHW", "NKD", "NKG", "NLG", "NNC", "NSC", "NT2", "NTL", "NVL", "NVN", "NVT",
	"OCB", "OGC", "OPC", "PAC", "PAN", "PC1", "PDN", "PDR", "PET", "PGC", "PGD", "PGI", "PHC", "PHR", "PHT", "PIT", "PJT", "PLP", "PLX", "PME", "PMG", "PNC", "PNJ", "POM", "POW", "PPC", "PSH", "PTB", "PTC", "PTL", "PVD", "PVF", "PVT", "PXI", "PXS", "PXT", "QBS", "QCG",
	"RAL", "RDP", "REE", "RIC", "ROS",
	"S4A", "SAB", "SAM", "SAV", "SBA", "SBC", "SBT", "SBV", "SC5", "SCD", "SCR", "SCS", "SEC", "SFC", "SFG", "SFI", "SGN", "SGR", "SGT", "SHA", "SHI", "SHP", "SII", "SJD", "SJF", "SJS", "SKG", "SMA", "SMB", "SMC", "SPM", "SRC", "SRF", "SSC", "SSI", "ST8", "STB", "STG", "STK", "SVC", "SVD", "SVI", "SVT", "SZC", "SZL",
	"TAC", "TBC", "TCB", "TCD", "TCH", "TCL", "TCM", "TCO", "TCR", "TCT", "TDC", "TDG", "TDH", "TDM", "TDP", "TDW", "TEG", "TGG", "THG", "THI", "TIC", "TIP", "TIX", "TLD", "TLG", "TLH", "TMP", "TMS", "TMT", "TN1", "TNA", "TNC", "TNH", "TNI", "TNT", "TPB", "TPC", "TRA", "TRC", "TRI", "TS4", "TSC", "TTA", "TTB", "TTE", "TTF", "TV2", "TVB", "TVS", "TVT", "TYA",
	"UDC", "UIC", "VAF", "VCB", "VCF", "VCG", "VCI", "VDP", "VDS", "VFG", "VGC", "VHC", "VHM", "VIB", "VIC", "VID", "VIP", "VIS", "VIX", "VJC", "VMD", "VND", "VNE", "VNG", "VNL", "VNM", "VNS", "VOS", "VPB", "VPD", "VPG", "VPH", "VPI", "VPL", "VPS", "VRC", "VRE", "VSC", "VSH", "VSI", "VTB", "VTF", "VTO", "YBM", "YEG",

	// HNX
	"AAV", "ACM", "ADC", "AGC", "ALT", "AMC", "AME", "AMV", "API", "APP", "APS", "ARM", "ART", "ATS", "AVS",
	"BAB", "BAX", "BBS", "BCC", "BCF", "BDB", "BED", "BII", "BKC", "BLF", "BNA", "BPC", "BSC", "BST", "BTS", "BTW", "BVS", "BXH",
	"C69", "C92", "CAG", "CAN", "CAP", "CDN", "CEO", "CET", "CIA", "CIC", "CJC", "CKV", "CLH", "CLM", "CMC", "CMS", "CPC", "CSC", "CTB", "CTC", "CTM", "CTP", "CTT", "CTV", "CTX", "CVN", "CX8",
	"D11", "DAD", "DAE", "DC2", "DDG", "DGL", "DHI", "DHL", "DHP", "DHT", "DIH", "DL1", "DNC", "DNM", "DNP", "DP3", "DPC", "DPS", "DS3", "DST", "DTD", "DTK", "DVG", "DXP", "DZM",
	"EBA", "EBS", "ECI", "EID", "EVS", "FDT", "FID", "FTV", "GBS", "GDW", "GFC", "GHA", "GIC", "GKM", "GLT", "GMA", "GMX",
	"HAD", "HAT", "HBB", "HBE", "HBS", "HCC", "HCT", "HDA", "HEV", "HGM", "HHC", "HHG", "HHL", "HJS", "HKB", "HKT", "HLC", "HLD", "HLY", "HMH", "HNM", "HOM", "HPC", "HPM", "HPR", "HPS", "HSC", "HST", "HTB", "HTC", "HTP", "HUT", "HVT",
	"ICG", "IDC", "IDJ", "IDV", "INC", "INN", "ITQ", "IVS", "KBT", "KDM", "KHS", "KKC", "KLF", "KLS", "KMF", "KMT", "KSD", "KSQ", "KST", "KTS", "KTT", "KVC",
	"L14", "L18", "L35", "L40", "L43", "L61", "L62", "LAS", "LBE", "LCD", "LCS", "LDP", "LHC", "LIG", "LM7", "LO5", "LUT",
	"MAC", "MAS", "MAX", "MBG", "MBS", "MCC", "MCF", "MCL", "MCO", "MDC", "MED", "MEL", "MHL", "MIH", "MIM", "MKV", "MNC", "MPT", "MSC", "MST", "MVB",
	"NAG", "NAP", "NBC", "NBP", "NBW", "NDN", "NDX", "NET", "NFC", "NGC", "NHC", "NHP", "NIS", "NLC", "NRC", "NSH", "NSN", "NST", "NTH", "NTP", "NVB", "NVC", "OCH", "ONE",
	"PBP", "PCE", "PCG", "PCT", "PDB", "PDC", "PEN", "PGN", "PGS", "PGT", "PHN", "PHP", "PIA", "PIC", "PJC", "PLC", "PMB", "PMC", "PMP", "PMS", "POT", "PPE", "PPP", "PPS", "PPY", "PRC", "PRE", "PSC", "PSD", "PSE", "PSI", "PSW", "PTD", "PTI", "PTS", "PV2", "PVB", "PVC", "PVG", "PVI", "PVL", "PVS",
	"QHD", "QST", "QTC", "RCL", "RHC",
	"S55", "S64", "S74", "S91", "S99", "SAF", "SCI", "SD2", "SD4", "SD5", "SD6", "SD9", "SDA", "SDC", "SDG", "SDN", "SDS", "SDT", "SDU", "SEB", "SED", "SEL", "SFN", "SGC", "SGD", "SGH", "SHB", "SHE", "SHN", "SHS", "SHT", "SIC", "SJ1", "SJC", "SJE", "SKS", "SLS", "SME", "SMN", "SMT", "SNG", "SPI", "SRA", "SSM", "SSS", "STC", "STP", "SVN", "SVS", "SZB",
	"TA9", "TAR", "TAS", "TBX", "TC6", "TCS", "TDN", "TDT", "TET", "TFC", "THB", "THD", "THS", "THT", "THV", "TIG", "TJC", "TKC", "TKU", "TLC", "TMB", "TMC", "TMX", "TNG", "TPH", "TPP", "TSB", "TSM", "TST", "TTC", "TTH", "TTL", "TTT", "TTZ", "TV3", "TV4", "TV6", "TVC", "TVD", "TXM", "UNI",
	"V12", "V21", "VAT", "VBC", "VC1", "VC2", "VC3", "VC6", "VC7", "VC9", "VCC", "VCH", "VCM", "VCS", "VCV", "VDL", "VE1", "VE2", "VE3", "VE4", "VE8", "VGP", "VGS", "VHE", "VHL", "VIE", "VIF", "VIG", "VIT", "VKC", "VLA", "VMC", "VMI", "VMS", "VNC", "VNF", "VNR", "VNT", "VSA", "VSM", "VTC", "VTH", "VTJ", "VTL", "VTS", "VTV", "VXB", "WCS", "WSS", "X20", "YSC",
	// UPCOM
	"ABC", "ABI", "ACV", "AFX", "AGX", "AMS", "APF", "AVC", "BAB", "BCM", "BDT", "BHA", "BLI", "BMS", "BNW", "BOT", "BPW", "BRR", "BSA", "BSG", "BSL", "BSQ", "BSR", "BT6", "BTC", "BUD", "BWS",
	"C21", "C4G", "CC1", "CDO", "CKD", "CLX", "CMF", "CMP", "CTR", "CTW", "DBW", "DDV", "DNA", "DNH", "DNN", "DNW", "DRI", "DSP", "DTK", "DVN", "DWS", "EIC", "EIN", "EVF", "EVS", "FIC", "FOC", "FOX",
	"G36", "GEG", "GHC", "GSM", "GTS", "GVR", "HAN", "HC3", "HDW", "HEM", "HEP", "HGW", "HIG", "HNA", "HND", "HNR", "HPI", "HPP", "HPW", "HRT", "HSM", "HTG", "HTM", "HUG", "HWS", "IDC", "IFS", "ILS", "IPA", "ISH",
	"KDF", "KHA", "KHW", "KLB", "KOS", "KSV", "KTC", "KTL", "LDW", "LIC", "LLM", "LPB", "LTG", "M10", "MCH", "MDF", "MIE", "MIG", "MKP", "MNB", "MPC", "MQB", "MSR", "MTA", "MVB", "MVC", "MVN",
	"NAW", "NBT", "NCP", "ND2", "NED", "NHH", "NLS", "NNG", "NQN", "NS2", "NTC", "OIL", "PDT", "PEG", "PFL", "PGV", "PIS", "POS", "PPH", "PRT", "PSB", "PSN", "PSP", "PVM", "PVP", "PXL", "QNS", "QTP",
	"RGC", "RTB", "SAS", "SBH", "SBM", "SCY", "SDI", "SEA", "SGP", "SID", "SJG", "SKH", "SNZ", "SQC", "SRT", "SSN", "SWC", "SZE", "TBD", "TCI", "TCW", "TDP", "THN", "TID", "TIS", "TL4", "TLP", "TMG", "TND", "TSJ", "TTP", "TVN",
	"VAV", "VCP", "VCW", "VEA", "VEC", "VEF", "VET", "VFC", "VGG", "VGI", "VGR", "VGT", "VGV", "VIB", "VIF", "VIN", "VIW", "VLB", "VLC", "VLW", "VNB", "VOC", "VSF", "VSN", "VTP", "WSB", "XDH",
}

func main() {
	// Get arguments
	symbol := flag.String("symbol", "", "Symbol that you want to get stats detail")
	daysBefore := flag.Int("day", 0, "0: today, -n: n days before")
	flag.Parse()

	// qs := &survey.Select{
	// 	Message: "Choose function:",
	// 	Options: []string{"Volume Price analyse"},
	// 	Default: "Volume Price analyse",
	// }

	// var fn string
	// survey.AskOne(qs, &fn)

	// if fn == "Volume Price analyse" {
	// 	priceVolumeAnalyse()
	// }

	// *symbol = "MBB"
	// *daysBefore = -19
	priceVolumeAnalyse(*symbol, *daysBefore)
}

func priceVolumeAnalyse(symbol string, daysBefore int) {
	if daysBefore > 0 {
		daysBefore = -1 * daysBefore
	}

	if len(symbol) == 0 {
		tm.Clear() // Clear current screen
		var sellList []string
		var buyList []string
		for _, input := range inputs {
			// remove the delimeter from the string
			input = strings.TrimSuffix(input, "\n")
			input = strings.ToUpper(input)
			pvStatsResult, err := stock.PriceVolumeStats(input, daysBefore)
			if err != nil {
				continue
			} else if pvStatsResult != nil {
				var isUpdated bool
				if pvStatsResult.Suggestion == "Buy" {
					isUpdated = true
					buyList = append(buyList, input)
					// fmt.Printf("\r\033[ABuy: %s\n", strings.Join(buyList, " "))

				} else if pvStatsResult.Suggestion == "Sell" {
					isUpdated = true
					sellList = append(sellList, input)
					// fmt.Printf("\rSell: %s", strings.Join(sellList, " "))
				}

				if isUpdated {
					tm.MoveCursor(1, 1)
					// Create Box with 30% width of current screen, and height of 20 lines
					buyBox := tm.NewBox(50|tm.PCT, 20, 5)
					sellBox := tm.NewBox(50|tm.PCT, 20, 5)

					// Add some content to the box
					// Note that you can add ANY content, even tables
					buyStr := fmt.Sprintf("Buy: %s", strings.Join(buyList, " "))
					sellStr := fmt.Sprintf("Sell: %s", strings.Join(sellList, " "))
					fmt.Fprint(buyBox, fmt.Sprintf("Date: %s\n\n%s", pvStatsResult.Date, util.WordWrap(buyStr, buyBox.Width-4)))
					fmt.Fprint(sellBox, fmt.Sprintf("Date: %s\n\n%s", pvStatsResult.Date, util.WordWrap(sellStr, sellBox.Width-4)))

					// Move Box to approx position of the screen
					tm.Print(tm.MoveTo(buyBox.String(), 0|tm.PCT, 0|tm.PCT))
					tm.Print(tm.MoveTo(sellBox.String(), 50|tm.PCT, 0|tm.PCT))

					tm.Flush()
				}
			}
		}
		fmt.Println()
	} else {
		symbol = strings.TrimSuffix(symbol, "\n")
		symbol = strings.ToUpper(symbol)
		pvStatsResult, err := stock.PriceVolumeStats(symbol, daysBefore)
		if err == nil {
			fmt.Println(pvStatsResult.ToString())
		} else {
			fmt.Println(err)
		}
	}
}

func main2() {
	// fmt.Print("Input ticker symbol: ")
	// reader := bufio.NewReader(os.Stdin)
	// // ReadString will block until the delimiter is entered
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }

	results := []stock.StatsResult{}

	f, err := os.Create("aobp_result.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, input := range inputs {

		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		input = strings.ToUpper(input)

		// TCBS
		translogDay, err := stock.TranslogStats(input)
		if err != nil {
			// fmt.Println(err)
			continue
		} else if translogDay != nil {
			if translogDay.TotalVol > 100000 {
				bidAskDay, err := stock.BidAskStats(input)
				if err != nil {
					fmt.Println(err)
					continue
				} else {
					if bidAskDay.OBPercent > 0.7 {
						bsaDay, err := stock.BSAStats(input)
						if err != nil {
							fmt.Println(err)
							continue
						} else {
							if bsaDay.Bsr < 0.7 {
								// Xa hang
								result := &stock.StatsResult{
									Time:          time.Now(),
									Ticker:        input,
									BuySellActive: bsaDay.Bsr,
									BidAskRatio:   bidAskDay.OBPercent,
									Volumn:        translogDay.TotalVol,
									Status:        "Xả",
									Suggestion:    "Bán",
								}
								fmt.Print(result.ToString())
								results = append(results, *result)
							}
						}
					} else if bidAskDay.OBPercent < 0.3 {
						bsaDay, err := stock.BSAStats(input)
						if err != nil {
							fmt.Println(err)
							continue
						} else {
							if bsaDay.Bsr > 1.3 {
								// Gom hang
								result := &stock.StatsResult{
									Time:          time.Now(),
									Ticker:        input,
									BuySellActive: bsaDay.Bsr,
									BidAskRatio:   bidAskDay.OBPercent,
									Volumn:        translogDay.TotalVol,
									Status:        "Gom",
									Suggestion:    "Mua",
								}
								fmt.Print(result.ToString())
								results = append(results, *result)
							}
						}
					}
				}
			}

			// _, err2 := f.WriteString(*result)

			// if err2 != nil {
			// 	log.Fatal(err2)
			// }

			// fmt.Println(*result)
			// time.Sleep(1 * time.Second)
		}
	}
}
